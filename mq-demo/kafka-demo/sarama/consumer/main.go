package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// Address kafka cluster addresses
var Address = []string{"117.51.148.112:9092"}

func main() {
	topic := []string{"test"}
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	// 广播式消费:消费者1
	go clusterConsumer(wg, Address, topic, "group-1")

	wg.Wait()
}

// 支持brokers cluster的消费者
func clusterConsumer(wg *sync.WaitGroup, brokers, topics []string, groupID string) {
	defer wg.Done()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// init consumer
	consumer, err := cluster.NewConsumer(brokers, groupID, topics, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupID, err)
		return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("%s:Error: %s\n", groupID, err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", groupID, ntf)
		}
	}()

	// consume messages, watch signals
	var successes int
Loop:
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", groupID, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
				successes++
			}
		case <-signals:
			break Loop
		}
	}
	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupID, successes)
}
