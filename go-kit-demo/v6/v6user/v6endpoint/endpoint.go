package v6endpoint

import (
	"context"

	"v6/v6user/pb"
	"v6/v6user/v6service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// EndPointServer ...
type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

// NewEndPointServer ...
func NewEndPointServer(servcie v6service.Service, log *zap.Logger, limit *rate.Limiter) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(servcie)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

// Login ...
func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

// MakeLoginEndPoint ...
func MakeLoginEndPoint(s v6service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (repsonse interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
