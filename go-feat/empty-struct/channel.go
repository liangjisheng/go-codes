package main

import (
	"fmt"
	"time"
)

//空结构体 与 channel 可谓是一个经典组合，有时候我们只是需要一个信号来控制程序的运行逻辑，并不在意其内容如何

func doTask1(ch chan struct{}) {
	time.Sleep(time.Second)
	fmt.Println("do task1")
	ch <- struct{}{}
}

func doTask2(ch chan struct{}) {
	time.Sleep(time.Second * 2)
	fmt.Println("do task2")
	ch <- struct{}{}
}

func Channel() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go doTask1(ch1)
	go doTask2(ch2)

	for {
		select {
		case <-ch1:
			fmt.Println("task1 done")
		case <-ch2:
			fmt.Println("task2 done")
		case <-time.After(time.Second * 5):
			fmt.Println("after 5 seconds")
			return
		}
	}
}
