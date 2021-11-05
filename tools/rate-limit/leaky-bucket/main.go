package main

import (
	"log"
	"math"
	"sync"
	"time"
)

//漏桶算法（Leaky Bucket），原理就是一个固定容量的漏桶，按照固定速率流出水滴
//漏桶限制的是常量流出速率（即流出速率是一个固定常量值），所以最大的速率就是出水的速率，不能应对出现突发流量

type LeakyBucket struct {
	rate float64 // 固定每秒出水速率
	capacity float64 // 桶的容量
	water float64 // 桶中当前水量
	lastLeakMs int64 // 桶上次漏水时间戳 ms
	lock sync.Mutex
}

func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.water = 0
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}

func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	eclipse := float64((now - l.lastLeakMs)) * l.rate / 1000 // 先执行漏水
	l.water = l.water - eclipse // 计算剩余水量
	l.water = math.Max(0, l.water) // 桶干了
	l.lastLeakMs = now

	if (l.water + 1) < l.capacity {
		// 尝试加水,并且水还未满
		l.water++
		return true
	} else {
		// 水满，拒绝加水
		return false
	}
}

func main() {
	var wg sync.WaitGroup
	var lr LeakyBucket
	lr.Set(1, 3) // 1s内最多请求3次
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
