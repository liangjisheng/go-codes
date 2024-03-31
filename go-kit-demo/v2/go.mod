module v2

go 1.13

require (
	github.com/go-kit/kit v0.10.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.14.0
	local.com/log-zap v0.0.0-00010101000000-000000000000
)

replace local.com/log-zap => ../log-zap
