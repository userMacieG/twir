package timers

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/lib/pq"
	"github.com/nicklaw5/helix/v2"
	"github.com/samber/lo"
	"github.com/satont/twir/apps/scheduler/internal/types"
	model "github.com/satont/twir/libs/gomodels"
	"github.com/satont/twir/libs/twitch"
	"go.uber.org/zap"
)

func NewStreams(ctx context.Context, services *types.Services) {
	timeTick := lo.If(services.Config.AppEnv != "production", 15*time.Second).Else(5 * time.Minute)
	ticker := time.NewTicker(timeTick)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				processStreams(services)
			}
		}
	}()
}

func processStreams(services *types.Services) {
	var channels []model.Channels
	err := services.Gorm.
		Where(`"isEnabled" = ? and "isBanned" = ?`, true, false).
		Select("id", `"isEnabled"`, `"isBanned"`).
		Find(&channels).Error
	if err != nil {
		zap.S().Error(err)
		return
	}

	usersIds := make([]string, len(channels))
	for i, channel := range channels {
		if !channel.IsEnabled && !channel.IsBanned {
			continue
		}

		usersIds[i] = channel.ID
	}

	discordIntegration := &model.Integrations{}
	err = services.Gorm.
		Where(`service = ?`, model.IntegrationServiceDiscord).
		Select("id").
		Find(discordIntegration).Error
	if err != nil {
		zap.S().Error(err)
		return
	}

	var discordIntegrations []model.ChannelsIntegrations
	if discordIntegration.ID != "" {
		err = services.Gorm.
			Where(`"integrationId" = ?`, discordIntegration.ID).
			Select("id", `"integrationId"`, "data").
			Find(&discordIntegrations).Error
		if err != nil {
			zap.S().Error(err)
			return
		}

		for _, integration := range discordIntegrations {
			if integration.Data == nil ||
				integration.Data.Discord == nil ||
				len(integration.Data.Discord.Guilds) == 0 {
				continue
			}

			for _, guild := range integration.Data.Discord.Guilds {
				if !guild.LiveNotificationEnabled {
					continue
				}
				usersIds = append(usersIds, guild.AdditionalUsersIdsForLiveCheck...)
			}
		}
	}

	usersIds = lo.Uniq(usersIds)

	var existedStreams []model.ChannelsStreams
	err = services.Gorm.Select("id", `"userId"`, `"parsedMessages"`).Find(&existedStreams).Error
	if err != nil {
		zap.S().Error(err)
		return
	}

	twitchClient, err := twitch.NewAppClient(*services.Config, services.Grpc.Tokens)
	if err != nil {
		zap.S().Error(err)
		return
	}

	chunks := lo.Chunk(usersIds, 100)
	wg := &sync.WaitGroup{}

	wg.Add(len(chunks))

	for _, chunk := range chunks {
		go func(chunk []string) {
			defer wg.Done()
			streams, err := twitchClient.GetStreams(
				&helix.StreamsParams{
					UserIDs: chunk,
				},
			)

			if err != nil || streams.ErrorMessage != "" {
				zap.S().Error(err)
				return
			}

			for _, userId := range chunk {
				twitchStream, twitchStreamExists := lo.Find(
					streams.Data.Streams, func(stream helix.Stream) bool {
						return stream.UserID == userId
					},
				)
				dbStream, dbStreamExists := lo.Find(
					existedStreams, func(stream model.ChannelsStreams) bool {
						return stream.UserId == userId
					},
				)

				tags := &pq.StringArray{}
				for _, tag := range twitchStream.Tags {
					*tags = append(*tags, tag)
				}

				channelStream := &model.ChannelsStreams{
					ID:             twitchStream.ID,
					UserId:         twitchStream.UserID,
					UserLogin:      twitchStream.UserLogin,
					UserName:       twitchStream.UserName,
					GameId:         twitchStream.GameID,
					GameName:       twitchStream.GameName,
					CommunityIds:   nil,
					Type:           twitchStream.Type,
					Title:          twitchStream.Title,
					ViewerCount:    twitchStream.ViewerCount,
					StartedAt:      twitchStream.StartedAt,
					Language:       twitchStream.Language,
					ThumbnailUrl:   twitchStream.ThumbnailURL,
					TagIds:         nil,
					Tags:           tags,
					IsMature:       twitchStream.IsMature,
					ParsedMessages: dbStream.ParsedMessages,
				}

				if twitchStreamExists && dbStreamExists {
					if result := services.Gorm.Where(
						`"userId" = ?`,
						userId,
					).Save(channelStream); result.Error != nil {
						zap.S().Error(
							result.Error,
							zap.String(
								"query", result.ToSQL(
									func(tx *gorm.DB) *gorm.DB {
										return tx.Where(`"userId" = ?`, userId).Save(channelStream)
									},
								),
							),
						)
						return
					}
				}

				if twitchStreamExists && !dbStreamExists {
					if result := services.Gorm.Where(
						`"userId" = ?`,
						userId,
					).Save(channelStream); result.Error != nil {
						zap.S().Error(
							result.Error,
							zap.String(
								"query", result.ToSQL(
									func(tx *gorm.DB) *gorm.DB {
										return tx.Where(`"userId" = ?`, userId).Save(channelStream)
									},
								),
							),
						)
						return
					}

					bytes, err := json.Marshal(
						&streamOnlineMessage{
							StreamID:  channelStream.ID,
							ChannelID: channelStream.UserId,
						},
					)
					if err != nil {
						zap.S().Error(err)
						return
					}

					services.PubSub.Publish("stream.online", bytes)
				}

				if !twitchStreamExists && dbStreamExists {
					err = services.Gorm.Where(
						`"userId" = ?`,
						userId,
					).Delete(&model.ChannelsStreams{}).Error
					if err != nil {
						zap.S().Error(err)
						return
					}

					bytes, err := json.Marshal(
						&streamOfflineMessage{
							ChannelID: channelStream.UserId,
						},
					)
					if err != nil {
						zap.S().Error(err)
						return
					}

					services.PubSub.Publish("stream.offline", bytes)
				}
			}
		}(chunk)
	}

	wg.Wait()
}

// { streamId: stream.id, channelId: channel }
type streamOnlineMessage struct {
	StreamID  string `json:"streamId"`
	ChannelID string `json:"channelId"`
}

type streamOfflineMessage struct {
	ChannelID string `json:"channelId"`
}
