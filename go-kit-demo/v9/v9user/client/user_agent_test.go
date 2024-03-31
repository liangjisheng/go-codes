package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"v9/utils"
	"v9/v9user/pb"
	"v9/v9user/v9service"

	"github.com/afex/hystrix-go/hystrix"
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
	hy := utils.NewHystrix("调用错误服务降级")
	cbs, _, _ := hystrix.GetCircuit("login")
	for i := 0; i < 1; i++ {
		time.Sleep(time.Millisecond * 100)
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		err = hy.Run("login", func() error {
			ack, err := userAgent.Login(context.Background(), &pb.Login{
				Account:  "liangjisheng",
				Password: "123456",
			})
			if err != nil {
				fmt.Println("err:", err)
				return err
			}
			fmt.Println(ack.Token)
			return nil
		})

		fmt.Println("熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		if err != nil {
			t.Log(err)
		}
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
	md := metadata.Pairs(v9service.ContextReqUUID, UUID)
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
