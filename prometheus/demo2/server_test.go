package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	for {
		time.Sleep(1000 * time.Millisecond)
		send()
	}
}

func send() {
	resp, err := http.Get("http://127.0.0.1:8080/index")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}
