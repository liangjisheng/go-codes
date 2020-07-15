package main

import (
	"fmt"
	"log"
	"net"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clientHandle(conn *net.UDPConn) {
	defer conn.Close()
	buf := make([]byte, 256)
	//读取数据
	//注意这里返回三个参数
	//第二个是udpaddr
	//下面向客户端写入数据时会用到
	_, udpaddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf))
	conn.WriteToUDP([]byte("hello,client \r\n"), udpaddr)
}

func main() {
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	chkError(err)
	//监听端口
	udpconn, err2 := net.ListenUDP("udp", udpaddr)
	chkError(err2)
	//udp没有对客户端连接的Accept函数
	for {
		clientHandle(udpconn)
	}
}
