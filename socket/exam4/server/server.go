package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "9999", "port")

type Msg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func main() {
	flag.Parse()
	var l net.Listener
	var err error

	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept() // 接收一个client
		if err != nil {
			fmt.Println("Error accepting:", err)
			os.Exit(1)
		}

		fmt.Printf("Received message %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected: " + ipStr)
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

		// 反序列化
		var msg Msg
		json.Unmarshal(b, &msg)
		fmt.Println("GET ==> data:", msg.Data, "type:", msg.Type)

		// 构建回复Msg
		resp := Resp{
			Data:   time.Now().String(),
			Status: 200,
		}
		r, _ := json.Marshal(resp)

		writer.Write(r)
		writer.Write([]byte("\n"))
		writer.Flush()
		//conn.Write(r)
		//conn.Write([]byte("\n"))
		fmt.Println("Done!")
	}
}
