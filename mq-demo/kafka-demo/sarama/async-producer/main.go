package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// kafka cluster addresses
var addresses = []string{"117.51.148.112:9092"}

func main() {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	// 注意 版本设置不对的话 kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	// 使用配置,新建一个异步生产者
	producer, err := sarama.NewAsyncProducer(addresses, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer producer.AsyncClose()

	// 循环判断哪个通道发送过来数据
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(1000 * time.Millisecond)
		value = "this is a message 0606 " + time.Now().Format("15:04:05")

		// 发送的消息 主题 注意: 这里的msg必须得是新构建的变量
		// 不然你会发现发送过去的消息内容都是一样的 因为批次发送消息的关系
		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.ByteEncoder(value),
		}

		producer.Input() <- msg // 使用通道发送
	}
}
