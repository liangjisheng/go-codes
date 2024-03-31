package v2service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// ContextReqUUID ...
const ContextReqUUID = "req_uuid"

// NewMiddlewareServer ...
type NewMiddlewareServer func(Service) Service

type logMiddlewareServer struct {
	logger *zap.Logger
	next   Service
}

// NewLogMiddlewareServer ...
func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return logMiddlewareServer{
			logger: log,
			next:   service,
		}
	}
}

func (l logMiddlewareServer) TestAdd(ctx context.Context, in Add) (out AddAck) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v2_service logMiddlewareServer", "TestAdd"), zap.Any("req", in), zap.Any("res", out))
	}()
	out = l.next.TestAdd(ctx, in)
	return out
}
