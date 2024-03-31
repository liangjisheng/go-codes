package v1endpoint

import (
	"context"
	"v1/v1service"

	"github.com/go-kit/kit/endpoint"
)

// EndPointServer ...
type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

// NewEndPointServer ....
func NewEndPointServer(svc v1service.Service) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}

// MakeAddEndPoint ...
func MakeAddEndPoint(s v1service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v1service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}
