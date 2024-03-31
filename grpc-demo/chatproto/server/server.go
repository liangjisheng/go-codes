package main

import (
	proto "go-demos/grpc-demo/chatproto/chat"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Streamer struct{}

func (s *Streamer) BidStream(stream proto.Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			request, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}
			// log.Println(request.Input)
			// log.Println(len(request.Input))

			// 如果接收正常，则根据接收到的字符串执行相应的指令
			switch request.Input {
			// 收到end指令，结束对话
			case "end":
				log.Println("对话结束")
				if err := stream.Send(&proto.Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到getData指令，连续返回 10 条数据
			case "getData":
				log.Println("收到返回数据流指令")
				for i := 0; i < 10; i++ {
					if err := stream.Send(&proto.Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}
			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到消息]: %s", request.Input)
				if err := stream.Send(&proto.Response{Output: "服务端返回: " + request.Input}); err != nil {
					return err
				}
			}
		}
	}
}

func main() {
	log.Println("启动服务器")
	server := grpc.NewServer()

	// 注册 ChatServer
	proto.RegisterChatServer(server, &Streamer{})
	address, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
