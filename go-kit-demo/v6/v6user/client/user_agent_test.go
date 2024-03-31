package client

import (
	"context"
	"os"
	"testing"
	"time"

	"v6/v6user/pb"
	"v6/v6user/v6service"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// go kit 客户端
func TestNewUserAgentClient(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := NewUserAgentClient([]string{"117.51.148.112:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		ack, err := userAgent.Login(context.Background(), &pb.Login{
			Account:  "alice",
			Password: "123456",
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(ack.Token)
	}
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
	md := metadata.Pairs(v6service.ContextReqUUID, UUID)
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
