package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

func main() {
	topic := "test"
	channel := "nsq_to_file"
	nsqAddr := "117.51.148.112:4150"

	go myconsumer(topic, channel, nsqAddr)
	go myconsumer(topic, channel, nsqAddr)

	select {}
}

// NSQHandler ...
type NSQHandler struct {
}

// HandleMessage ...
func (*NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func myconsumer(topic, channel, nsqAddr string) {
	config := nsq.NewConfig()
	config.MaxInFlight = 9
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if nil != err {
		fmt.Println("err", err)
		return
	}

	consumer.AddHandler(&NSQHandler{})
	err = consumer.ConnectToNSQD(nsqAddr)
	if nil != err {
		fmt.Println("err", err)
		return
	}
}
