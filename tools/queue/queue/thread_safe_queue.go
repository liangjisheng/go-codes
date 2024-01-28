package queue

import "sync"

type NonBlockingQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewBlockingQueue(n int) (q *NonBlockingQueue) {
	return &NonBlockingQueue{data: make([]interface{}, 0, n)}
}

func (q *NonBlockingQueue) Add(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

func (q *NonBlockingQueue) Poll() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
