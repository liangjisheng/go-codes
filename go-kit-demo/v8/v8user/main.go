package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"v8/utils"
	"v8/v8user/pb"
	"v8/v8user/v8endpoint"
	"v8/v8user/v8service"
	"v8/v8user/v8transport"

	"github.com/go-kit/kit/log"
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

var grpcAddr = flag.String("g", "127.0.0.1:8080", "grpcAddr")
var prometheusAddr = flag.String("p", "127.0.0.1:8081", "prometheus addr")
var quitChan = make(chan error, 1)

func main() {
	flag.Parse()
	var (
		etcdAddrs      = []string{"117.51.148.112:2379"}
		serName        = "svc.user.agent"
		grpcAddr       = *grpcAddr
		prometheusAddr = *prometheusAddr
		ttl            = 5 * time.Second
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
		count := metricsprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "Number of requests",
		}, []string{"method"})

		histogram := metricsprometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "Request consumes time",
		}, []string{"method"})

		golangLimit := rate.NewLimiter(10, 1)
		server := v8service.NewService(utils.GetLogger(), count, histogram)
		endpoints := v8endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit)
		grpcServer := v8transport.NewGRPCServer(endpoints, utils.GetLogger())
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			utils.GetLogger().Warn("[user_agent] Listen", zap.Error(err))
			quitChan <- err
		}
		registrar.Register()
		utils.GetLogger().Info("[user_agent] grpc run " + grpcAddr)
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
		pb.RegisterUserServer(baseServer, grpcServer)
		if err = baseServer.Serve(grpcListener); err != nil {
			utils.GetLogger().Warn("[user_agent] Serve", zap.Error(err))
			quitChan <- err
			return
		}
	}()

	go func() {
		utils.GetLogger().Info("[user_agent] prometheus run " + prometheusAddr)
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(prometheusAddr, m)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-quitChan
	registrar.Deregister()
	utils.GetLogger().Info("[user_agent] quit err", zap.Error(err))
}
