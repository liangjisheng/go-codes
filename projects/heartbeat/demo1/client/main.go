package main

import (
	"net"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

func main() {
	server := "127.0.0.1:8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		beego.Warning("Fatal error:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		beego.Warning("Fatal error:", err.Error())
		os.Exit(1)
	}

	generator(conn)
	beego.Trace("send over")
}

func generator(conn *net.TCPConn) {
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " times" + "--> heart beating from client"
		num, err := conn.Write([]byte(words))
		if err != nil {
			beego.Warning("Fatal error:", err.Error())
			os.Exit(1)
		}
		beego.Trace("Server have receive:", num, " bytes")
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 2; i++ {
		beego.Info("client stop 1s")
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " times" + "--> heart beating from client"
		num, err := conn.Write([]byte(words))
		if err != nil {
			beego.Warning("Fatal error:", err.Error())
			os.Exit(1)
		}
		beego.Trace("Server have receive:", num, " bytes")
		time.Sleep(1 * time.Second)
	}
}
