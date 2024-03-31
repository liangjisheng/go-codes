package v5service

import (
	"context"
	"errors"
	"fmt"

	"v5/utils"
	"v5/v5user/pb"

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
func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v5_service Service", "Login 处理请求"))
	if in.Account != "alice" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v5_service Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}
