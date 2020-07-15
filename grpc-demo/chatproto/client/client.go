package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"

	proto "go-demos/grpc-demo/chatproto/chat"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败: [%v]", err)
		return
	}
	defer conn.Close()

	client := proto.NewChatClient(conn)
	ctx := context.Background()
	// 创建双向流数据
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Printf("创建数据流失败: [%v]", err)
	}

	// 启动一个 goroutine 接收命令行输入的指令
	go func() {
		log.Println("请输入消息...")
		input := bufio.NewReader(os.Stdin)
		for {
			// 获取 命令行输入的字符串， 以回车 作为结束标志
			str, _ := input.ReadString('\n')
			if containLine := strings.Contains(str, "\n"); containLine {
				str = string([]rune(str)[:len(str)-2])
			}
			if err := stream.Send(&proto.Request{Input: str}); err != nil {
				return
			}
		}
	}()

	for {
		// 接收从服务端返回的数据流
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务器的结束信号")
			break // 跳出循环，结束客户端程序
		}
		if err != nil {
			log.Println("接收数据错误:", err)
		}
		log.Printf("[客户端收到]: %s", response.Output)
	}
}
