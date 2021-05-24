package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doTimeoutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println("deadline set:", deadline.Unix())
			if time.Now().After(deadline) {
				log.Println("error:", ctx.Err())
			}
		}

		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("working")
		}
	}
}

func demo1() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	fmt.Println("current time:", time.Now().Unix())
	fmt.Println("timeout: 2 seconds")
	go doTimeoutStuff(ctx)
	time.Sleep(3 * time.Second)
}
