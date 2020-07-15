package main

import (
	"context"
	"go-demos/grpc-demo/streamDemo/pro"
	"log"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
)

const (
	ADDRESS = "localhost:50051"
)

func main() {
	// 通过grpc库建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	// 通过刚刚的连接生成一个client对象
	c := pro.NewGreeterClient(conn)
	// 调用服务端推送流
	reqStreamData := &pro.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqStreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
	// 客户端推送流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pro.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			if data != nil {
				log.Println(data)
			} else {
				log.Println("Recv over")
				break
			}
		}
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamReqData{Data: "ssss"})
			time.Sleep(time.Second)
		}
	}()

	select {}
}
