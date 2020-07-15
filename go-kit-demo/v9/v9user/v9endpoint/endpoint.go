package v9endpoint

import (
	"context"
	"fmt"

	"v9/v9user/pb"
	"v9/v9user/v9service"

	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// EndPointServer ...
type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

// NewEndPointServer ...
func NewEndPointServer(servcie v9service.Service, log *zap.Logger, limit *rate.Limiter, tracer opentracing.Tracer) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(servcie)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		// loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
		loginEndPoint = NewTracerEndpointMiddleware(tracer)(loginEndPoint)
	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

// Login ...
func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		fmt.Println("s.LoginEndPoint", err)
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

// MakeLoginEndPoint ...
func MakeLoginEndPoint(s v9service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (repsonse interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
