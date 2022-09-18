package main

import (
	"fmt"
	"sync"
)

//Lock/Unlock 必须成对出现。一直不调用 Unlock 会导致死锁，对未加锁的 Mutex 调用 Unlock 会导致 panic
//Mutex 不可复用。sync 的同步原语都不能复制使用（比如作为参数传入），因为可能不是初始状态（state 标记）

type safeInt struct {
	sync.Mutex
	Num int
}

func mutex() {
	count := safeInt{}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			fmt.Print(count.Num, " ")
			count.Unlock()
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
