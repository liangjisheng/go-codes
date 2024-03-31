package client

import (
	"context"
	"testing"

	"v5/v5user/pb"
	"v5/v5user/v5service"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	logzap "local.com/log-zap"
)

// go kit 客户端
func TestGrpcClient(t *testing.T) {
	logger := logzap.NewLogger(
		logzap.SetAppName("go-kit"),
		logzap.SetDevelopment(true),
		logzap.SetLevel(zap.DebugLevel),
	)
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	svr := NewGRPCClient(conn, logger)
	ack, err := svr.Login(context.Background(), &pb.Login{
		Account:  "alice",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ack.Token)
}

// grpc 原生客户端
func TestGrpc(t *testing.T) {
	serviceAddress := "127.0.0.1:8080"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	userClient := pb.NewUserClient(conn)
	UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
	md := metadata.Pairs(v5service.ContextReqUUID, UUID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := userClient.RPCUserLogin(ctx, &pb.Login{
		Account:  "liangjisheng",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res.Token)
}
