package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu Mutex

	// 启动 goroutine，在一段时间持有锁。
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)

	// 尝试获取到锁
	if mu.TryLock() {
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}

	// 没有获取到
	fmt.Println("can't get the lock")
}

func TestMutex(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n",
		mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
