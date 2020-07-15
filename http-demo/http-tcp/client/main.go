package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("server is not starting")
		return
	}
	defer conn.Close()

	for {
		fmt.Print("input send data:")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			continue
		} else {
			fmt.Printf("client send: %s", input)
		}

		b := []byte(input)
		conn.Write(b) // write in connect
	}
}
