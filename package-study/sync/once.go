package main

import (
	"fmt"
	"sync"
)

// Once 的作用是多次调用但只执行一次，Once 只有一个方法，Once.Do()，
// 向 Do 传入一个函数，这个函数在第一次执行 Once.Do() 的时候会被调用，
// 以后再执行 Once.Do() 将没有任何动作，即使传入了其它的函数，也不会被执行，
// 如果要执行其它函数，需要重新创建一个 Once 对象

func once() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // 多次调用只执行一次
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
