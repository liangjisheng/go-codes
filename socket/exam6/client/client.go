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

func main() {
	//获取udpaddr
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	chkError(err)
	//连接，返回udpconn
	udpconn, err2 := net.DialUDP("udp", nil, udpaddr)
	chkError(err2)
	//写入数据
	_, err3 := udpconn.Write([]byte("client\r\n"))
	chkError(err3)
	buf := make([]byte, 256)
	//读取服务端发送的数据
	_, err4 := udpconn.Read(buf)
	chkError(err4)
	fmt.Println(string(buf))
}
