package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

// 垃圾回收一直是Go语言的一块心病, 在它执行垃圾回收的时间中, 你很难做什么
// 在垃圾回收压力大的服务中, GC占据的CPU有可能超过2%, 造成的Pause经常超过2ms
// 垃圾严重的时候，秒级的GC也出现过
// 如果经常临时使用一些大型结构体, 可以用Pool来减少GC

// sync.Pool是一个可以存或取的临时对象集合
// sync.Pool可以安全被多个线程同时使用, 保证线程安全
// 注意、注意、注意, sync.Pool 中保存的任何项都可能随时不做通知的释放掉
// 所以不适合用于像socket长连接或数据库连接池
// sync.Pool主要用途是增加临时对象的重用率, 减少GC负担

// 1. 缓存对象没有数量限制 即只受制于物理限制 - 内存
// 2. 缓存对象的过期  注册了  runtime_registerPoolCleanup(poolCleanup) 函数,  每次gc之前都会调用.
// sync.Pool的缓存的期限只是在两次gc之间...  因此不能实现socket连接池.

type structR6 struct {
	B1 [100000]int
}

var r6Pool = sync.Pool{
	New: func() interface{} {
		return new(structR6)
	},
}

func withPool() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		sr6 := r6Pool.Get().(*structR6)
		sr6.B1[0] = 0
		r6Pool.Put(sr6)
	}
	fmt.Println("pool Used:", time.Since(startTime))
}

func standard() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		var sr6 structR6
		sr6.B1[0] = 0
	}
	fmt.Println("standard Used:", time.Since(startTime))
}

var (
	pool1 = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
)

func withoutPool1() {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		buf := &bytes.Buffer{}
		buf.WriteString("hello")
		buf.WriteString("world")
	}
	fmt.Println("Without pool:", time.Since(start))
}

func withPool1() {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		buf := pool1.Get().(*bytes.Buffer)
		buf.WriteString("hello")
		buf.WriteString("world")
		pool1.Put(buf)
	}

	fmt.Println("With pool:", time.Since(start))
}
