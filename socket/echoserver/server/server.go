package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleRequest(conn net.Conn) {
	log.Printf("Accepting new connection %v", conn.RemoteAddr())
	close := func() {
		log.Println("Closing connection")
		conn.Close()
	}
	defer close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("Client close connection")
			break
		} else if err != nil {
			log.Println("read message form client err:", err)
			break
		}

		str = strings.TrimSpace(str)
		if str == "STOP" {
			log.Println("Receive stop signal")
			break
		}

		writer.WriteString(fmt.Sprintf("> %s\n", str))
		writer.Flush()
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("err:", err)
		os.Exit(-1)
	}
	log.Println("Listen on port 8080")
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("new connect err:", err)
			continue
		}

		go handleRequest(conn)
	}
}
