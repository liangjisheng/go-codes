package simple_worker_pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job ...
type Job struct {
	id        int
	randomNum int
}

// Result ...
type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 工作函数
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker2(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomNum)}
		results <- output
	}
	wg.Done() // 引用计数减1
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker2(&wg)
	}
	wg.Wait()
	close(results) // 所有协程结束后，关闭输出结果信道
}

// 分配工作协程
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNum := rand.Intn(999)
		job := Job{i, randomNum}
		jobs <- job
	}
	close(jobs) // 所有工作分配完后，关闭输入工作信道
}

// 读取结果协程
func result(done chan bool) {
	for r := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n",
			r.job.id, r.job.randomNum, r.sumOfDigits)
	}
	done <- true // 所有结果读取完成后，给主协程发个通知
}
