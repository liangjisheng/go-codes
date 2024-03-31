package main

import (
	"log"
	"sync"
	"time"
)

func cond4() {
	var mailbox uint8 // 信箱，0 表示空，1 表示满
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock) // 参数为 lock 的指针值
	recvCond := sync.NewCond(&lock)

	signal := make(chan struct{}, 3)
	max := 3

	// 发信
	go func(max int) {
		defer func() {
			signal <- struct{}{}
		}()

		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()        // 锁住信箱
			for mailbox == 1 { // 信箱满
				sendCond.Wait() // 等待信箱空
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1 // 发信，信箱满
			log.Printf("sender [%d]: the mailbox has been sent.", i)
			lock.Unlock()     // 解锁信箱
			recvCond.Signal() // 通知接受者
		}
	}(max)

	// 收信
	go func(max int) {
		defer func() {
			signal <- struct{}{}
		}()

		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 500)
			// lock.RLock() // panic
			lock.Lock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", i)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", i)
			// lock.RUnlock()
			lock.Unlock()
			sendCond.Signal()
		}
	}(max)

	<-signal
	<-signal
}
