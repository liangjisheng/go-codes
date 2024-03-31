package main

import (
	"fmt"
	"net"
)

//使用了同步阻塞IO函数来进行IO操作，所以程序会在执行IO操作时被阻塞。这意味着，在等待 IO 操作完成的同时
//程序无法执行其他任务。因此，在处理多个客户端连接时，我们使用了 goroutine 来进行并发处理。

func main() {
	// 监听TCP端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("TCP server started and listening on port 8080")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting client connection:", err)
			continue
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的数据
	//程序会在执行IO操作时被阻塞
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error reading data from client:", err)
		return
	}

	// 处理客户端发送的数据
	response := processData(data)

	// 向客户端发送响应数据
	//程序会在执行IO操作时被阻塞
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response to client:", err)
		return
	}
}

func processData(data []byte) string {
	// 处理客户端发送的数据
	return "Processed data: " + string(data)
}
