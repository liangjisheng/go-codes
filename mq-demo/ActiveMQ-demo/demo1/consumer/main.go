package main

import (
	"fmt"
	"time"

	"github.com/go-stomp/stomp"
)

func recvData(ch chan *stomp.Message) {
	for {
		v := <-ch
		fmt.Println(string(v.Body))
	}
}

func main() {
	ch := make(chan *stomp.Message)
	go recvData(ch)

	conn, err := stomp.Dial("tcp", "117.51.148.112:61613")
	if err != nil {
		fmt.Println("dial err =", err)
		return
	}

	sub, err := conn.Subscribe("testQ", stomp.AckMode(stomp.AckAuto))
	for {
		select {
		case v := <-sub.C:
			ch <- v
		case <-time.After(time.Second * 30):
			return
		}
	}
}
