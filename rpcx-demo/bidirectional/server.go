package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server address")
	clientConn net.Conn
	connected  = false
)

// Arith ...
type Arith int

// Mul ...
func (t *Arith) Mul(ctx context.Context, args *service.Args, reply *service.Reply) error {
	clientConn = ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	connected = true
	return nil
}

func main() {
	flag.Parse()

	ln, _ := net.Listen("tcp", "127.0.0.1:9981")
	go http.Serve(ln, nil)

	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(Arith), "")
	go s.Serve("tcp", *addr)

	for !connected {
		time.Sleep(time.Second)
	}

	fmt.Printf("start to send messages to %s\n", clientConn.RemoteAddr().String())
	for {
		if clientConn != nil {
			err := s.SendMessage(clientConn, "test_service_path", "test_service_method", nil, []byte("abcde"))
			if err != nil {
				fmt.Printf("failed to send messsage to %s: %v\n", clientConn.RemoteAddr().String(), err)
				clientConn = nil
			}
		}
		time.Sleep(time.Second)
	}
}
