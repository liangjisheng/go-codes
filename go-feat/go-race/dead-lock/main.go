package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		for i := 0; i < 2; i++ {
			c <- "hello there"
		}
		//如果不关闭 chan，则下面的 read 会一直等待，造成死锁
		//close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

//$ go run main.go
//$ go run -race main.go
