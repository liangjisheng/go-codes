package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"v6/utils"
	"v6/v6user/pb"
	"v6/v6user/v6endpoint"
	"v6/v6user/v6service"
	"v6/v6user/v6transport"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

var grpcAddr = flag.String("g", "127.0.0.1:8080", "grpcAddr")
var quitChan = make(chan error, 1)

func main() {
	flag.Parse()
	var (
		etcdAddrs = []string{"117.51.148.112:2379"}
		serName   = "svc.user.agent"
		grpcAddr  = *grpcAddr
		ttl       = 5 * time.Second
	)

	utils.NewLoggerServer()

	// 初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		utils.GetLogger().Error("[user_agent]  NewClient", zap.Error(err))
		return
	}
	registrar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%s", serName, grpcAddr),
		Value: grpcAddr,
	}, log.NewNopLogger())

	go func() {
		golangLimit := rate.NewLimiter(1, 10)
		server := v6service.NewService(utils.GetLogger())
		endpoints := v6endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit)
		grpcServer := v6transport.NewGRPCServer(endpoints, utils.GetLogger())
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			utils.GetLogger().Warn("[user_agent] Listen", zap.Error(err))
			quitChan <- err
		}
		registrar.Register()
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
		pb.RegisterUserServer(baseServer, grpcServer)
		if err = baseServer.Serve(grpcListener); err != nil {
			utils.GetLogger().Warn("[user_agent] Serve", zap.Error(err))
			quitChan <- err
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	utils.GetLogger().Info("[user_agent] run " + grpcAddr)
	err = <-quitChan
	registrar.Deregister()
	utils.GetLogger().Info("[user_agent] quit err", zap.Error(err))
}
