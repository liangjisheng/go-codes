package main

import (
	"fmt"
	"go-demos/grpc-demo/streamDemo/pro"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50051"

type server struct {
}

//服务端 单向流
func (s *server) GetStream(req *pro.StreamReqData, res pro.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&pro.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

//客户端 单向流
func (s *server) PutStream(cliStr pro.Greeter_PutStreamServer) error {
	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err:", err)
			break
		}
	}

	return nil
}

//客户端服务端 双向流
func (s *server) AllStream(allStr pro.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
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
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamResData{Data: "ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT) // 监听端口
	if err != nil {
		return
	}

	s := grpc.NewServer()                   // 创建一个grpc服务器
	pro.RegisterGreeterServer(s, &server{}) // 注册事件
	s.Serve(lis)                            // 处理链接
}
