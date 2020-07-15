package v3endpoint

import (
	"context"

	"v3/v3service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// EndPointServer ...
type EndPointServer struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

// NewEndPointServer ...
func NewEndPointServer(svc v3service.Service, log *zap.Logger) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
	}
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint, LoginEndPoint: loginEndPoint}
}

// MakeAddEndPoint ...
func MakeAddEndPoint(s v3service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v3service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

// MakeLoginEndPoint ...
func MakeLoginEndPoint(s v3service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v3service.Login)
		return s.Login(ctx, req)
	}
}
