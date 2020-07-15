package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("dial connect err:", err)
		os.Exit(-1)
	}
	defer conn.Close()

	// make some noise
	go func() {
		for {
			conn.Write([]byte("hello\n"))
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		buf := make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("%v\n", err)
		}
		fmt.Printf("%s %v\n", string(buf), n)
	}
}
