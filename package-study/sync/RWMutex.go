package main

import (
	"sync"
	"time"
)

//RWMutex 是基于 Mutex 实现的，其对写操作优先，即如果已有 writer 在等待请求锁（阻塞的 Lock 调用），
//则会阻止新的 reader 获取锁，优先保障 writer（注意不是抢占）

type CounterRW struct {
	mu    sync.RWMutex
	count uint64
}

func (c *CounterRW) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *CounterRW) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func RWMutex1() {
	var counter CounterRW
	for i := 0; i < 10; i++ {
		go func() {
			for {
				counter.Count() // 读频率 1次/ms
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for {
		counter.Incr() // 写频率 1次/s
		time.Sleep(time.Second)
	}
}
