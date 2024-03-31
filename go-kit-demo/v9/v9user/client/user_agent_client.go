package client

import (
	"context"
	"io"
	"time"
	"v9/utils"

	"v9/v9user/pb"
	"v9/v9user/v9endpoint"
	"v9/v9user/v9service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gofrs/uuid"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UserAgent ...
type UserAgent struct {
	instancerm *etcdv3.Instancer
	logger     log.Logger
	tracer     opentracing.Tracer
}

// NewUserAgentClient ...
func NewUserAgentClient(addr []string, logger log.Logger) (*UserAgent, error) {
	var (
		etcdAddrs = addr
		serName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	utils.NewLoggerServer()
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	tracer, _, err := utils.NewJaegerTracer("user_agent_client")
	if err != nil {
		return nil, err
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		return nil, err
	}
	instancerm, err := etcdv3.NewInstancer(etcdClient, serName, logger)
	if err != nil {
		return nil, err
	}
	return &UserAgent{
		instancerm: instancerm,
		logger:     logger,
		tracer:     tracer,
	}, err
}

// UserAgentClient ...
func (u *UserAgent) UserAgentClient() (v9service.Service, error) {
	var (
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)
	var endpoints v9endpoint.EndPointServer
	{
		factory := u.factoryFor(v9endpoint.MakeLoginEndPoint)
		endpointer := sd.NewEndpointer(u.instancerm, factory, u.logger)
		balancer := lb.NewRandom(endpointer, time.Now().UnixNano())
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.LoginEndPoint = retry
	}
	return endpoints, nil
}

func (u *UserAgent) factoryFor(makeEndpoint func(v9service.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		chainUnaryServer := grpcmiddleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(u.tracer)),
			grpc_zap.UnaryClientInterceptor(utils.GetLogger()),
			// utils.JaegerServerMiddleware(tracer),
		)
		conn, err := grpc.Dial(instance, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chainUnaryServer))
		if err != nil {
			return nil, nil, err
		}
		srv := u.NewGRPCClient(conn)
		endpoints := makeEndpoint(srv)
		return endpoints, conn, err
	}
}

// NewGRPCClient ...
func (u *UserAgent) NewGRPCClient(conn *grpc.ClientConn) v9service.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
			md.Set(v9service.ContextReqUUID, UUID)
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
	return v9endpoint.EndPointServer{
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
