package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server has been start===>")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8000")

	// 服务器一般不定位具体的客户端套接字
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	ConnMap := make(map[string]*net.TCPConn)
	for {
		tcpConn, _ := tcpListener.AcceptTCP()
		defer tcpConn.Close()

		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("连接的客户端信息:", tcpConn.RemoteAddr().String())
	}
}
