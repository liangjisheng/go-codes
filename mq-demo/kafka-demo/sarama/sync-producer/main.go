package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// kafka cluster addresses
var addresses = []string{"117.51.148.112:9092"}

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 1 * time.Second
	p, err := sarama.NewSyncProducer(addresses, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()

	topic := "test"
	srcValue := "sync: this is a message. index=%d"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功 partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(1 * time.Second)
	}
}
