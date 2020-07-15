package v4endpoint

import (
	"context"

	"v4/v4service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// EndPointServer ...
type EndPointServer struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

// NewEndPointServer ...
func NewEndPointServer(svc v4service.Service, log *zap.Logger, limit *rate.Limiter, limiter ratelimit.Limiter) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
		addEndPoint = NewUberRateMiddleware(limiter)(addEndPoint)
	}
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint, LoginEndPoint: loginEndPoint}
}

// MakeAddEndPoint ...
func MakeAddEndPoint(s v4service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v4service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

// MakeLoginEndPoint ...
func MakeLoginEndPoint(s v4service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v4service.Login)
		return s.Login(ctx, req)
	}
}
