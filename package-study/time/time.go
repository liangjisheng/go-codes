package main

import (
	"fmt"
	"time"
)

func time1() {
	// 一定时间后，自动塞一个当前时间到channel中
	tchan := time.After(time.Second * 3)
	fmt.Printf("tchan type=%T\n", tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=", <-tchan)
	fmt.Println("mark 2")
}

func time2() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	// 这个分支首先执行，然后主协程退出，上面go启动的协程并不会执行
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}

func time3() {
	c := make(chan int)
	timeout := time.After(time.Second * 2)
	t1 := time.NewTimer(time.Second * 3) // 效果相同，只执行一次
	var i int

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C:
				fmt.Println("3s定时任务")
			case <-timeout:
				i++
				fmt.Println(i, "2s定时输出")
			// case <- time.After(time.Second)  :
			// 是本次监听动作的超时时间， 意思就说，只有在本次select 操作中会有效，
			// 再次select 又会重新开始计时（从当前时间+4秒后）， 但是有default ，
			// 那case 超时操作，肯定执行不到了
			case <-time.After(time.Second * 4):
				fmt.Println("4s timeout...")
			default:
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 6)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}

func sender(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		if i >= 5 {
			time.Sleep(time.Second * 7)
		} else {
			time.Sleep(time.Second)
		}
	}
}

func time4() {
	c := make(chan int)
	go sender(c)
	timeout := time.After(time.Second * 3)
	for {
		select {
		case d := <-c:
			fmt.Println(d)
		case <-timeout:
			fmt.Println("这是定时操作任务 >>>>>")
		case dd := <-time.After(time.Second * 3):
			fmt.Println(dd, "这是超时*****")
		}

		fmt.Println("for end")
	}
}
