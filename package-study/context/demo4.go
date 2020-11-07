package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func demo4() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(&r)

	// 监控程序打印前，需要检测 r.Context() 是否已经结束
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <- r.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in progress")
			}
		}
	}()

	time.Sleep(2*time.Second)
	w.Write([]byte("hi"))
}