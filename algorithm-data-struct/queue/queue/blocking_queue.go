package queue

import (
	"fmt"
	"strings"
	"sync"
)

type BlockingQueue struct {
	cond     *sync.Cond
	data     []interface{}
	capacity int
	logs     []string
}

func NewQueue(capacity int) *BlockingQueue {
	return &BlockingQueue{
		cond:     &sync.Cond{L: &sync.Mutex{}},
		data:     make([]interface{}, 0),
		capacity: capacity, logs: make([]string, 0),
	}
}

func (q *BlockingQueue) Add(d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.data) == q.capacity {
		q.cond.Wait()
	}
	q.data = append(q.data, d)
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("En %v\n", d))
	// 通知其他 waiter 进行 Poll 或 Add 操作
	q.cond.Broadcast()

}

func (q *BlockingQueue) Poll() (d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == 0 {
		q.cond.Wait()
	}
	d = q.data[0]

	q.data = q.data[1:]

	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("De %v\n", d))

	// 通知其他 waiter 进行 Poll 或 Add 操作
	q.cond.Broadcast()
	return
}

func (q *BlockingQueue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.data)
}

func (q *BlockingQueue) String() string {
	var b strings.Builder
	for _, log := range q.logs {
		//fmt.Fprint(&b, log)
		b.WriteString(log)
	}
	return b.String()
}
