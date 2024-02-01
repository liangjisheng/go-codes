package main

//timer 是高性能定时器库
//timer利用时间轮算法，通过降低定时器精度的方式，将同一个时间单位内的任务集中存储到一个双向链表，
//可以一次锁操作处理，减少锁竞争，进而提高性能，对于业务中有大量定时任务，同时对精度要求大于10ms
//的场景，可以尝试timer库来优化。

import (
	"log"
	"time"

	"github.com/antlabs/timer"
)

// 一次性定时器
func demo1() {
	tm := timer.NewTimer()

	tm.AfterFunc(1*time.Second, func() {
		log.Printf("after\n")
	})

	tm.AfterFunc(2*time.Second, func() {
		log.Printf("after\n")
	})
	tm.Run()
}

// 周期性定时器
func demo2() {
	tm := timer.NewTimer()

	tm.ScheduleFunc(1*time.Second, func() {
		log.Printf("schedule\n")
	})

	tm.Run()
}

type customTest struct {
	count int
}

// Next 只要实现Next接口就行
func (c *customTest) Next(now time.Time) (rv time.Time) {
	rv = now.Add(time.Duration(c.count) * time.Millisecond * 10)
	c.count++
	return
}

// 自定义周期性定时器
func demo3() {
	tm := timer.NewTimer(timer.WithMinHeap())
	_ = tm.CustomFunc(&customTest{count: 1}, func() {
		log.Printf("%v\n", time.Now())
	})
	tm.Run()
}

// 取消某一个定时器
func demo4() {
	tm := timer.NewTimer()

	// 只会打印2 time.Second
	tm.AfterFunc(2*time.Second, func() {
		log.Printf("2 time.Second")
	})

	// tk3 会被 tk3.Stop()函数调用取消掉
	tk3 := tm.AfterFunc(3*time.Second, func() {
		log.Printf("3 time.Second")
	})

	tk3.Stop() //取消tk3

	tm.Run()
}

func main() {
	//demo1()
	//demo2()
	//demo3()
	demo4()
}
