package client

import (
	"context"
	"v5/v5user/v5endpoint"

	"v5/v5user/pb"
	"v5/v5user/v5service"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// NewGRPCClient ...
func NewGRPCClient(conn *grpc.ClientConn, log *zap.Logger) v5service.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			md.Set(v5service.ContextReqUUID, UUID)
			ctx = metadata.NewOutgoingContext(context.Background(), *md)
			return ctx
		}),
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"pb.User",
			"RPCUserLogin",
			RequestLogin,
			ResponseLogin,
			pb.LoginAck{},
			options...).Endpoint()
	}
	return v5endpoint.EndPointServer{
		LoginEndPoint: loginEndpoint,
	}
}

// RequestLogin ...
func RequestLogin(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Login)
	return &pb.Login{Account: req.Account, Password: req.Password}, nil
}

// ResponseLogin ...
func ResponseLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginAck)
	return &pb.LoginAck{Token: resp.Token}, nil
}
