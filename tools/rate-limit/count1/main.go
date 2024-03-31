package main

import (
	"log"
	"sync"
	"time"
)

//计数器是一种最简单限流算法，其原理就是：在一段时间间隔内，对请求进行计数，与阀值进行比较判断是否需要限流，一旦到了时间临界点，将计数器清零
//可以在程序中设置一个变量 count，当过来一个请求我就将这个数+1，同时记录请求时间。
//当下一个请求来的时候判断 count 的计数值是否超过设定的频次，以及当前请求的时间和第一次请求时间是否在 1 分钟内。
//如果在 1 分钟内并且超过设定的频次则证明请求过多，后面的请求就拒绝掉。
//如果该请求与第一个请求的间隔时间大于计数周期，且 count 值还在限流范围内，就重置 count

type Counter struct {
	rate int // 计数周期内最多允许的请求数
	begin time.Time // 计数开始时间
	cycle time.Duration // 计数周期
	count int // 计数周期内累计收到的请求数
	lock sync.Mutex
}

func (l *Counter) Set(r int, cycle time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

func (l *Counter) Reset(t time.Time) {
	l.count = 0
	l.begin = t
}

func (l *Counter) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count == l.rate - 1 {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			// 速度允许范围内, 重置计数器
			l.Reset(now)
			return true
		} else {
			return false // 限速
		}
	} else {
		// 没有达到速率限制，计数加1
		l.count++
		return true
	}
}

func main() {
	var wg sync.WaitGroup
	var lr Counter
	lr.Set(3, time.Second) // 1s内最多请求3次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		log.Println("创建请求:", i)
		go func(i int) {
			if lr.Allow() {
				log.Println("响应请求:", i)
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
