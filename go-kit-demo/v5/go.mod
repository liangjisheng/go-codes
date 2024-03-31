module v5

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.5
	go.uber.org/zap v1.14.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.28.0
	local.com/log-zap v0.0.0-00010101000000-000000000000
)

replace local.com/log-zap => ../log-zap
