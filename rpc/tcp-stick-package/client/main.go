package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"time"
)

// User ...
type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Msg  string `json:"msg"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panicln(err)
	}

	for {
		msg := getRandString()
		u := User{
			Name: "alice",
			Age:  24,
			Msg:  msg,
		}
		data, err := json.Marshal(u)
		if err != nil {
			log.Panicln(err)
		}
		dataLen := len(data)
		b := make([]byte, dataLen+12)
		b[0] = 'V'
		b[1] = '1'
		binary.BigEndian.PutUint64(b[2:10], uint64(time.Now().Unix()))
		binary.BigEndian.PutUint16(b[10:12], uint16(dataLen))
		copy(b[12:], data)
		_, err = conn.Write(b)
		if err != nil {
			log.Panicln(err)
		}
		time.Sleep(time.Second)
	}
}

func getRandString() string {
	length := rand.Intn(50)
	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = byte(rand.Intn(26) + 97)
	}
	return string(strBytes)
}
