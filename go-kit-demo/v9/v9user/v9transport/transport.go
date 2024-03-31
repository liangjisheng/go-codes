package v9transport

import (
	"context"
	"fmt"

	"v9/v9user/pb"
	"v9/v9user/v9endpoint"
	"v9/v9user/v9service"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	login grpctransport.Handler
}

// NewGRPCServer ...
func NewGRPCServer(endpoint v9endpoint.EndPointServer, log *zap.Logger) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, v9service.ContextReqUUID, md.Get(v9service.ContextReqUUID))
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
