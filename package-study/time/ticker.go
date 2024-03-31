package main

import (
	"fmt"
	"time"
)

// type Ticker //主要用来按照指定的时间周期来调用函数或者计算表达式，
// 通常的使用方式是利用go新开一个协程使用，它是一个断续器

func ticker() {
	// 新生成一个ticker,此Ticker包含一个channel，此channel以给定的duration发送时间
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 2500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
