package main

import (
	"fmt"
)

func main() {
	can := make(chan int, 8)
	can <- 1
	can <- 1
	can <- 1

	close(can)

	/*	for {
			select {
			case tem, ok := <-can:
				if ok == false {
					// select 一个nil的channel会直接返回
					return
				} else {
					fmt.Println(tem)
				}
				time.Sleep(time.Second)
			}
		}
	*/

	// 如果该channle没有被关闭，会一直阻塞，如果关闭了，会自动退出循环
	for i := range can {
		fmt.Println(i)
	}
}
