package v2endpoint

import (
	"context"
	"v2/v2service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// EndPointServer ...
type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

// NewEndPointServer ...
func NewEndPointServer(svc v2service.Service, log *zap.Logger) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}

// MakeAddEndPoint ...
func MakeAddEndPoint(s v2service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v2service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}
