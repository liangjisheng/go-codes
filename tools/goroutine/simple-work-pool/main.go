package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

// 创建2个信道，messages 用于传送任务消息，result 用于接收消息处理结果
// 创建3个 Worker 协程，用于接收和处理来自 messages 信道的任务消息，并将处理结果通过信道 result 返回
// 通过信道 messages 发布10条任务
// 通过信道 result 接收任务处理结果

// Message ...
type Message struct {
	ID   int
	Name string
}

func main() {
	messages := make(chan Message, 100)
	result := make(chan error, 100)

	// 创建任务处理Worker
	for i := 0; i < 3; i++ {
		go worker(i, messages, result)
	}

	total := 0
	// 发布任务
	for k := 1; k <= 10; k++ {
		messages <- Message{ID: k, Name: "job" + strconv.Itoa(k)}
		total++
	}

	close(messages)

	// 接受任务处理结果
	for j := 1; j <= total; j++ {
		res := <-result
		if res != nil {
			log.Println(res.Error())
		}
	}

	close(result)
}

func worker(worker int, msg <-chan Message, result chan<- error) {
	// 从通道 chan Message 中监听&接收新的任务
	for job := range msg {
		log.Println("worker:", worker, "msg: ", job.ID, ":", job.Name)
		// 模拟任务执行时间
		time.Sleep(time.Second * time.Duration(RandInt(1, 3)))
		// 通过通道返回执行结果
		result <- nil
	}
}

// RandInt ...
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
