package main

// 模拟一个基于HTTP协议的客户端请求去连接一个Web服务端

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	// ResolveTCPAddr获取一个TCPAddr
	// net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)。
	// addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError("ResolveTCPAddr", err)

	// 与远程主机tcpAddr建立tcp连接
	// net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
	// laddr表示本机地址，一般设置为nil
	// raddr表示远程的服务地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError("DialTCP", err)

	// 将内容写入socket,并发送请求信息
	// writer := bufio.NewWriter(conn)
	// writer.Write([]byte("HEAD / HTTP/1.0\n"))
	// writer.Flush()
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\n"))
	checkError("Write", err)

	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	checkError("ReadLine", err)
	fmt.Println(string(line))

	// 不能用 ioutil.ReadAll(conn),不知道为啥
	// 最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息
	// result, err := ioutil.ReadAll(conn)
	// checkError("ReadAll", err)
	// fmt.Println(string(result))
	os.Exit(0)
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s Fatal error: %s", str, err.Error())
	}
}
