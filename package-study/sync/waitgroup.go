package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup用于等待一批Go协程执行结束，程序控制会一直阻塞，知道这些协程全部执行完毕

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("start Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // 引用计数减1
}

func waitgroup() {
	no := 3
	// wg是一个结构体，使用引用计数来工作
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)          // 引用计数加1
		go process(i, &wg) // 必须使用指针，否则会产生wg的一个拷贝
	}
	wg.Wait() // 等待引用计数变为0，解除主协程的阻塞
	fmt.Println("All go routines finished executing")
}
