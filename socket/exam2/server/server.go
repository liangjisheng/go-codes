package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := "127.0.0.1:57777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 必须先读取conn里面的内容，然后才能写入
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		checkError(err)
		fmt.Println(string(b[:n]))

		// reader := bufio.NewReader(conn)
		// line, _, err := reader.ReadLine()
		// checkError(err)
		// fmt.Println(string(line))

		// 上面两种读取方式都能读取conn里的内容
		// 但不能用 ioutil.ReadAll(conn),不知道为啥
		// 使用ioutil.ReadAll(conn)目前来看是有问题的,程序在这里停下来了
		// result, err := ioutil.ReadAll(conn)
		// checkError(err)
		// fmt.Println(string(result))

		// go handleRequest(conn)
		daytime := time.Now().String()
		writer := bufio.NewWriter(conn)
		writer.Write([]byte(daytime))
		writer.Write([]byte("\n"))
		writer.Flush()
		// conn.Write([]byte(daytime)) // don't care about return value
		conn.Close() // we're finished with this client
	}
}

func handleRequest(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected :" + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// 读取一行数据, 以"\n"结尾
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(b))

		writer.Write([]byte("server response\n"))
		writer.Flush()
		//conn.Write(r)
		//conn.Write([]byte("\n"))
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
