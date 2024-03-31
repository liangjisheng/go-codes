package queue

import (
	"testing"
	"time"
)

func TestBlockingQueue(t *testing.T) {
	q := NewQueue(5)

	go func() {
		res := q.Poll()
		t.Log("G1", res)
	}()

	time.Sleep(100 * time.Millisecond)
	go func() {
		res := q.Poll()
		t.Log("G2", res)
	}()

	time.Sleep(100 * time.Millisecond)
	q.Add(1)
	time.Sleep(100 * time.Millisecond)
	q.Add(2)
	time.Sleep(100 * time.Millisecond)
}

func TestLockFreeQueue(t *testing.T) {
	q := NewLKQueue()

	go func() {
		for {
			v, ok := q.Poll().(int)
			if !ok {
				continue
			}
			t.Log(v)
		}
	}()

	time.Sleep(200 * time.Millisecond)
	q.Add(1)
	time.Sleep(200 * time.Millisecond)
	q.Add(1)
	time.Sleep(200 * time.Millisecond)
}
