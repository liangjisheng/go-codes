module v9

go 1.13

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.3.0
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.etcd.io/etcd v3.3.0+incompatible
	go.uber.org/zap v1.14.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.28.0
	local.com/log-zap v0.0.0-00010101000000-000000000000
)

replace local.com/log-zap => ../log-zap
