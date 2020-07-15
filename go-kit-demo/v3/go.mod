module v3

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/gofrs/uuid v3.2.0+incompatible
	go.uber.org/zap v1.14.0
	local.com/log-zap v0.0.0-00010101000000-000000000000
)

replace local.com/log-zap => ../log-zap
