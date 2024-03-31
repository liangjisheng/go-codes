package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	go http.ListenAndServe("127.0.0.1:8080", nil)

	server, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func handleConn(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
}
