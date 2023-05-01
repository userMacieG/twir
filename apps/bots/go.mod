module github.com/satont/tsuwari/apps/bots

go 1.20

require (
	github.com/aidenwallis/go-ratelimiting v0.0.1
	github.com/gempir/go-twitch-irc/v3 v3.2.0
	github.com/getsentry/sentry-go v0.18.0
	github.com/imroc/req/v3 v3.32.0
	github.com/nicklaw5/helix/v2 v2.22.1
	github.com/prometheus/client_golang v1.14.0
	github.com/redis/go-redis/v9 v9.0.2
	github.com/samber/do v1.6.0
	github.com/samber/lo v1.37.0
	github.com/satont/tsuwari/libs/config v0.0.0-00010101000000-000000000000
	github.com/satont/tsuwari/libs/gomodels v0.0.0-20221114143619-e5e207524b96
	github.com/satont/tsuwari/libs/gopool v0.0.0-00010101000000-000000000000
	github.com/satont/tsuwari/libs/grpc v0.0.0-00010101000000-000000000000
	github.com/satont/tsuwari/libs/pubsub v0.0.0-00010101000000-000000000000
	github.com/satont/tsuwari/libs/twitch v0.0.0-00010101000000-000000000000
	github.com/satont/tsuwari/libs/types v0.0.0-20230501072144-eb4ca703ed7c
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/postgres v1.4.8
	gorm.io/gorm v1.24.6
)

require (
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v9 v9.0.0-rc.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	github.com/guregu/null v4.0.0+incompatible // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/onsi/ginkgo/v2 v2.8.3 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-18 v0.2.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.2.1 // indirect
	github.com/quic-go/qtls-go1-20 v0.1.1 // indirect
	github.com/quic-go/quic-go v0.32.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20230213192124-5e25df0256eb // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

replace github.com/satont/tsuwari/libs/config => ../../libs/config

replace github.com/satont/tsuwari/libs/gomodels => ../../libs/gomodels

replace github.com/satont/tsuwari/libs/twitch => ../../libs/twitch

replace github.com/satont/tsuwari/libs/grpc => ../../libs/grpc

replace github.com/satont/tsuwari/libs/pubsub => ../../libs/pubsub

replace github.com/satont/tsuwari/libs/gopool => ../../libs/gopool
