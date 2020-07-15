package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"grpclbbalancer/etcdv3"
	pb "grpclbbalancer/proto"
)

var (
	// EtcdEndpoints etcd地址
	EtcdEndpoints = []string{"117.51.148.112:2379"}
	// SerName 服务名称
	SerName    = "simple_grpc"
	grpcClient pb.SimpleClient
)

func main() {
	r := etcdv3.NewServiceDiscovery(EtcdEndpoints)
	resolver.Register(r)
	// 连接服务器
	target := fmt.Sprintf("%s:///%s", r.Scheme(), SerName)
	conn, err := grpc.Dial(
		target,
		grpc.WithBalancerName("weight"),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	for i := 0; i < 50; i++ {
		time.Sleep(1 * time.Second)
		route(i)
	}
}

// route 调用服务端Route方法
func route(i int) {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc " + strconv.Itoa(i),
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
