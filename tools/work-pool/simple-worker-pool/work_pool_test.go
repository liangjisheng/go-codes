package simple_worker_pool

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestWorkerPool1(t *testing.T) {
	messages := make(chan Message, 100)
	result := make(chan error, 100)

	// 创建任务处理Worker
	for i := 0; i < 3; i++ {
		go worker1(i, messages, result)
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

func TestWorkerPool2(t *testing.T) {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done // 主协程等待读取结果协程将所有结果读取完毕

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
