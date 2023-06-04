package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//用原子操作来替换mutex锁,其主要原因是，原子操作由底层硬件支持
//而锁则由操作系统提供的API实现。若实现相同的功能，前者通常会更有效率

type Counter interface {
	Inc()
	Load() int64
}

//CommonCounter 普通版
type CommonCounter struct {
	counter int64
}

func (c *CommonCounter) Inc() {
	c.counter++
}

func (c *CommonCounter) Load() int64 {
	return c.counter
}

//MutexCounter 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

//AtomicCounter 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func dataCount(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}
