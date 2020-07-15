package main

// 在这个协议中我们使用前2个字节表示协议版本,8个字节表示时间戳,
// 2个字节表示数据长度,前12字节是属于协议头部部分,剩下的是数据部分

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	// ProtocolVersion ...
	ProtocolVersion = "V1"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()
	header := make([]byte, 12)
	for {
		// read packet header
		_, err := conn.Read(header)
		if err != nil {
			return
		}
		if fmt.Sprintf("%s", header[:2]) != ProtocolVersion {
			log.Println("valid protoc version")
			return
		}

		timestamp := binary.BigEndian.Uint64(header[2:10])
		t := time.Unix(int64(timestamp), 0).Format("2006-01-02 03:04:05 PM")
		log.Printf("client send data time %s", t)

		length := int16(binary.BigEndian.Uint16(header[10:]))
		log.Println("data length", length)

		// read data
		databuf := make([]byte, length)
		_, err = conn.Read(databuf)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", databuf)
		conn.Write(databuf)
	}
}
