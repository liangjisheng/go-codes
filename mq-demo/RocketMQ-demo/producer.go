package main

import (
	"fmt"

	rocketmq "github.com/apache/rocketmq-client-go/core"
)

func sendMessage(config *rocketmq.ProducerConfig) {
	producer, err := rocketmq.NewProducer(config)
	if err != nil {
		fmt.Println("create common producer failed, error:", err)
		return
	}

	err = producer.Start()
	if err != nil {
		fmt.Println("start common producer error", err)
		return
	}
	defer producer.Shutdown()

	fmt.Printf("Common producer: %s started... \n", producer)
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%s-%d", "Hello,Common MQ Message-", i)
		result, err := producer.SendMessageSync(&rocketmq.Message{Topic: "TestTopic", Body: msg})
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("send message: %s result: %s\n", msg, result)
	}
	fmt.Println("shutdown common producer.")
}

func main0() {
	pConfig := &rocketmq.ProducerConfig{
		ClientConfig: rocketmq.ClientConfig{
			GroupID:    "GID_0",
			NameServer: "http://117.51.148.112:9876",
		},
		// Set to Common Producer as default.
		ProducerModel: rocketmq.CommonProducer,
	}
	sendMessage(pConfig)
}
