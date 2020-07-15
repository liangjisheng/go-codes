package v7service

import (
	"context"
	"fmt"

	"v7/v7user/pb"

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

func (l logMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("err", err))
	}()
	out, err = l.next.Login(ctx, in)
	return
}
