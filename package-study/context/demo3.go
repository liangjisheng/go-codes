package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func testA(ctx context.Context) {
	ctxA, _ := context.WithTimeout(ctx, (50 * time.Millisecond))
	ch := make(chan int)
	go testB(ctxA, ch)

	select {
	case <-ctx.Done():
		fmt.Println("testA Done")
		return
	case i := <-ch:
		fmt.Println(i)
	}
}

func testB(ctx context.Context, ch chan int) {
	sumCh := make(chan int)
	go func(sumCh chan int) {
		sum := 10
		time.Sleep(100 * time.Millisecond)
		sumCh <- sum
	}(sumCh)

	select {
	case <-ctx.Done():
		fmt.Println("testB Done")
		<-sumCh
		return
	case i := <-sumCh:
		fmt.Println("send", i)
		ch <- i
	}
}

func demo3() {
	go http.ListenAndServe(":8080", nil)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	go testA(ctx)
	select {}
}
