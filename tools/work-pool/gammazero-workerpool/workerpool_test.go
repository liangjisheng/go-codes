package gammazero_workerpool

//Concurrency limiting goroutine pool. Limits the concurrency of task execution,
//not the number of tasks queued. Never blocks submitting tasks, no matter how
//many tasks are queued.

import (
	"fmt"
	"testing"
	"time"

	"github.com/gammazero/workerpool"
)

func TestSimple(t *testing.T) {
	wp := workerpool.New(2)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		v := r
		wp.Submit(func() {
			time.Sleep(time.Second)
			fmt.Println("Handling request:", v)
		})
	}

	wp.StopWait()
}
