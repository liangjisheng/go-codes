package simple_worker_pool

import (
	"log"
	"math/rand"
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

func worker1(worker int, msg <-chan Message, result chan<- error) {
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
	return min + rand.Intn(max-min+1)
}
