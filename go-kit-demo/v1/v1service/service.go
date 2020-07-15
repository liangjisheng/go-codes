package v1service

import "context"

// Service ...
type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
}

// NewService ...
func NewService() Service {
	return &baseServer{}
}

func (s baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	return AddAck{Res: in.A + in.B}
}
