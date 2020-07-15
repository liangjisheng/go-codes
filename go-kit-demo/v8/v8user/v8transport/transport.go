package v8transport

import (
	"context"
	"fmt"

	"v8/v8user/pb"
	"v8/v8user/v8endpoint"
	"v8/v8user/v8service"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	login grpctransport.Handler
}

// NewGRPCServer ...
func NewGRPCServer(endpoint v8endpoint.EndPointServer, log *zap.Logger) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, v8service.ContextReqUUID, md.Get(v8service.ContextReqUUID))
			return ctx
		}),
		grpctransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
	}

	return &grpcServer{
		login: grpctransport.NewServer(
			endpoint.LoginEndPoint,
			RequestGrpcLogin,
			ResponseGrpcLogin,
			options...,
		),
	}
}

// RPCUserLogin ...
func (s *grpcServer) RPCUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		fmt.Println("s.login.ServeGRPC", err)
		return nil, err
	}
	return rep.(*pb.LoginAck), nil
}

// RequestGrpcLogin ...
func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Login)
	return &pb.Login{Account: req.GetAccount(), Password: req.GetPassword()}, nil
}

// ResponseGrpcLogin ...
func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginAck)
	return &pb.LoginAck{Token: resp.Token}, nil
}
