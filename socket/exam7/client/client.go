package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	recvBytes, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(recvBytes))
}
