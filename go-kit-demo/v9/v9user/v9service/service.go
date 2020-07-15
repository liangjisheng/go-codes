package v9service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"v9/utils"
	"v9/v9user/pb"

	"github.com/go-kit/kit/metrics"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Service ...
type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

// NewService ...
func NewService(log *zap.Logger, counter metrics.Counter, histogram metrics.Histogram, tracer opentracing.Tracer) Service {
	var server Service
	server = &baseServer{log}
	server = NewTracerMiddlewareServer(tracer)(server)
	server = NewLogMiddlewareServer(log)(server)
	server = NewMetricsMiddlewareServer(counter, histogram)(server)
	return server
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v9_service Service", "Login 处理请求"))
	if in.Account != "liangjisheng" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}

	// 模拟耗时
	// rand.Seed(time.Now().UnixNano())
	// sl := rand.Int31n(10-1) + 1
	// time.Sleep(time.Duration(sl) * time.Millisecond * 100)
	// 模拟错误 熔断那个版本用到
	// if rand.Intn(10) > 3 {
	// 	err = errors.New("服务器运行错误")
	// 	return
	// }

	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("调用 v9_service Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}
