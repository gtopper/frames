module github.com/v3io/frames

go 1.12

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9
	github.com/golang/protobuf v1.2.0
	github.com/nuclio/errors v0.0.1
	github.com/nuclio/logger v0.0.1
	github.com/nuclio/nuclio-test-go v0.0.0-20180704132150-0ce6587f8e37 // indirect
	github.com/nuclio/zap v0.0.2
	github.com/pavius/impi v0.0.0-20200212064320-5db7efa5f87b // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/v3io/v3io-go v0.1.5-0.20200413162202-5d20cf2c5c71
	github.com/v3io/v3io-tsdb v0.9.24
	github.com/valyala/fasthttp v1.2.0
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	google.golang.org/grpc v1.17.0
)

replace (
	github.com/valyala/fasthttp => github.com/gtopper/fasthttp v0.0.0-00000000000000-7784c88efdc6142d543c33811c84a0c62ab3c4e7
	github.com/xwb1989/sqlparser => github.com/v3io/sqlparser v0.0.0-20190306105200-4d7273501871
)
