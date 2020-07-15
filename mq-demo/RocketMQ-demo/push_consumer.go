package main

import (
	"fmt"
	"sync/atomic"

	rocketmq "github.com/apache/rocketmq-client-go/core"
)

func consumeWithPush(config *rocketmq.PushConsumerConfig) {
	consumer, err := rocketmq.NewPushConsumer(config)
	if err != nil {
		println("create Consumer failed, error:", err)
		return
	}

	ch := make(chan interface{})
	var count = (int64)(1000000)
	// ********************************************
	// MUST subscribe topic before consumer started.
	// *********************************************
	consumer.Subscribe("TestTopic", "*", func(msg *rocketmq.MessageExt) rocketmq.ConsumeStatus {
		fmt.Printf("A message received, MessageID:%s, Body:%s \n", msg.MessageID, msg.Body)
		if atomic.AddInt64(&count, -1) <= 0 {
			ch <- "quit"
		}
		return rocketmq.ConsumeSuccess
	})

	err = consumer.Start()
	if err != nil {
		println("consumer start failed,", err)
		return
	}

	fmt.Printf("consumer: %s started...\n", consumer)
	<-ch
	err = consumer.Shutdown()
	if err != nil {
		println("consumer shutdown failed")
		return
	}
	println("consumer has shutdown.")
}

func main1() {
	pConfig := &rocketmq.PushConsumerConfig{
		ClientConfig: rocketmq.ClientConfig{
			GroupID:    "GID_0",
			NameServer: "http://117.51.148.112:9876",
		},
		Model:         rocketmq.Clustering,
		ConsumerModel: rocketmq.CoCurrently,
	}
	consumeWithPush(pConfig)
}
