package main

import (
	"context"
	"fmt"
	"time"
)

func demo5() {
	ctx,cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	go Speak(ctx)
	time.Sleep(10*time.Second)
}

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}