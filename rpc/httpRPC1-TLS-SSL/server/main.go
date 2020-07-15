package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

// Result ...
type Result struct {
	Num, Ans int
}

// Cal ...
type Cal int

// Square ...
func (cal *Cal) Square(num int, res *Result) error {
	res.Num = num
	res.Ans = num * num
	return nil
}

func main() {
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", "127.0.0.1:8080", config)
	log.Printf("Serving RPC server on port %d", 8080)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
