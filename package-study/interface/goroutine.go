package main

import (
	"fmt"
	"runtime"
)

// 一段耗时的计算函数
func consumer(ch chan int) {
	for {
		data := <-ch
		if 0 == data {
			break
		}

		fmt.Println(data)
	}

	fmt.Println("goroutine exit")
}

func main() {
	ch := make(chan int)

	for {
		var dummy string
		// 获取输入, 模拟进程持续运行
		fmt.Scan(&dummy)

		if "quit" == dummy {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}

			continue
		}

		go consumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
