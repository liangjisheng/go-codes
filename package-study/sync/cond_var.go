// 条件变量可以协调想要访问共享资源的线程
// 当共享资源的状态发生变化时，它可以被用来通知被互斥锁阻塞的线程

// 条件变量的 Wait 方法做了什么
// 1. 把调用它的 goroutine 加入到当前条件变量的通知队列中
// 2. 解锁当前的条件变量基于的那个互斥锁。所以在调用该 Wait 方法前，必须先锁定互斥锁，
// 否则会引发不可恢复的 panic
// 3. 让当前 goroutine 处于等待状态，等到通知到来时再决定是否唤醒它
// 此时，该 goroutine 就会阻塞在调用该 Wait 方法的那行代码上
// 4. 如果通知到来并决定唤醒该 goroutine，则在唤醒它后重写锁定当前条件变量基于的互斥锁
// 自此后，当前的 goroutine 就会继续执行后面的代码了

// 使用 for 检测资源状态的原因
// 如果一个 goroutine 因收到通知而被唤醒，但却发现共享资源的状态依然不符合它的要求
// 那么就应该再次调用条件变量的 Wait 方法，并继续等待下次通知的到来

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cond1() {
	condition := false // 条件不满足
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	go func() {
		mu.Lock()
		condition = true // 更改条件
		cond.Signal()    // 发送通知：条件已经满足
		mu.Unlock()
	}()

	mu.Lock()
	// 检查条件是否满足，避免虚假通知，同时避免 Signal 提前于 Wait 执行
	for !condition {
		// 等待条件满足的通知，如果收到虚假通知，则循环继续等待
		cond.Wait() // 等待时 mu 处于解锁状态，唤醒时重新锁定
	}

	fmt.Println("条件满足，开始后续动作...")
	mu.Unlock()
}

func cond2() {
	runtime.GOMAXPROCS(4)
	testCond()
}

func testCond() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		time.Sleep(time.Second * 1)
		c.L.Lock()
		fmt.Println("[1] 变更condition状态,并发出变更通知.")
		condition = true
		c.Signal() // c.Broadcast()
		fmt.Println("[1] 继续后续处理.")
		c.L.Unlock()
	}()

	c.L.Lock()
	fmt.Println("[2] condition..........1")
	for !condition {
		fmt.Println("[2] condition..........2")
		// 等待Cond消息通知
		c.Wait()
		fmt.Println("[2] condition..........3")
	}
	fmt.Println("[2] condition..........4")
	c.L.Unlock()

	fmt.Println("main end...")
}
