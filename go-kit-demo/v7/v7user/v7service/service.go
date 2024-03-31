package v7service

import (
	"context"
	"errors"
	"fmt"

	"v7/utils"
	"v7/v7user/pb"

	"github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
)

// Service ...
type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

// NewService ...
func NewService(log *zap.Logger, counter metrics.Counter, histogram metrics.Histogram) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	server = NewMetricsMiddlewareServer(counter, histogram)(server)
	return server
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v7_service Service", "Login 处理请求"))
	if in.Account != "liangjisheng" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v7_service Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}
