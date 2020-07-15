package main

import "time"

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()

	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
			// simulate	模仿
			time.Sleep(time.Second)
		}
	}()

	return out
}

func main() {
	in := producer(10)
	ch := square(in)

	// consumer
	for _ = range ch {

	}
}
