package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "watch exit.")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine watching...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func demo2() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "liangjisheng")
	go watch(valueCtx)
	time.Sleep(2 * time.Second)
	fmt.Println("ok, notify watch exit.")
	cancelFunc()

	time.Sleep(3 * time.Second)
}

func demo2_1() {
	ctx := context.WithValue(context.Background(), key, "alice")
	v := ctx.Value("name")
	fmt.Println(v)
	v = ctx.Value("name1")
	fmt.Println(v)
}
