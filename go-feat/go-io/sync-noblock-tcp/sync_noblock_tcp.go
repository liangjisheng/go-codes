package main

import (
	"fmt"
	"net"
	"time"
)

//在"handleConnection"函数中，我们首先设置了读取和写入超时时间，以防止IO操作阻塞程序。

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

	// 设置读取超时时间为5秒
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// 读取客户端发送的数据
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error reading data from client:", err)
		return
	}

	// 处理客户端发送的数据
	response := processData(data[:n])

	// 设置写入超时时间为5秒
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	// 向客户端发送响应数据
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
