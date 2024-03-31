package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	nsqAddr := "117.51.148.112:4150"
	producer, err := nsq.NewProducer(nsqAddr, nsq.NewConfig())
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		msg := "message: " + time.Now().Format("2006-01-02 15:04:05")
		err = producer.Publish("test", []byte(msg))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("send:", msg)
		time.Sleep(1 * time.Second)
	}
}
