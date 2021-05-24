package main

import (
	"context"
	"fmt"
	"time"
)

func myPrintln(ctx context.Context, a, b int) {
	for {
		fmt.Println(a + b)
		a, b = a+1, b+1
		select {
		case <-ctx.Done():
			fmt.Println("program over.")
			return
		}
	}
}

func demo() {
	a, b := 1, 2
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Microsecond)
		cancelFunc()
	}()
	myPrintln(ctx, a, b)

	time.Sleep(20 * time.Microsecond)
}
