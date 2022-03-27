package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//用原子操作来替换mutex锁,其主要原因是，原子操作由底层硬件支持
//而锁则由操作系统提供的API实现。若实现相同的功能，前者通常会更有效率

func atomic1() {
	var a, b int32 = 13, 13
	var c int32 = 9
	res := atomic.CompareAndSwapInt32(&a, b, c)
	fmt.Println("swapped:", res)
	fmt.Println("替换的值:", c)
	fmt.Println("替换之后a的值:", a)

	fmt.Println()
	a, b = 13, 12
	old := atomic.SwapInt32(&a, b)
	fmt.Println("old的值:", old)
	fmt.Println("替换之后a的值", a)

	fmt.Println()
	addValue := atomic.AddInt32(&a, 1)
	fmt.Println("增加之后:", addValue)
	delValue := atomic.AddInt32(&a, -4)
	fmt.Println("减少之后:", delValue)
}

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

func test(c Counter) {
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

func atomic2() {
	c1 := CommonCounter{} // 非并发安全
	test(&c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}
