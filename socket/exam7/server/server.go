package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen 函数会监听来自 9090 端口的连接，返回一个 net.Listener 对象
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Accept 函数会阻塞程序，直到接收到来自端口的连接
		// 每接收到一个链接，就会返回一个 net.Conn 对象表示这个连接
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintln(conn, "Hello form TCP server")
		conn.Close()
	}
}
