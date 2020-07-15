package main

import (
	"net"
	"os"

	"v5/utils"
	"v5/v5user/pb"
	"v5/v5user/v5endpoint"
	"v5/v5user/v5service"
	"v5/v5user/v5transport"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(1, 10)
	server := v5service.NewService(utils.GetLogger())
	endpoints := v5endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit)
	grpcServer := v5transport.NewGRPCServer(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 127.0.0.1:8080")
	grpcListener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		utils.GetLogger().Warn("Listen", zap.Error(err))
		os.Exit(0)
	}
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		utils.GetLogger().Warn("Serve", zap.Error(err))
		os.Exit(0)
	}
}
