package main

import (
	"fmt"
	"sync"
)

type safeInt struct {
	sync.Mutex
	Num int
}

func mutex() {
	count := safeInt{}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			fmt.Print(count.Num, " ")
			count.Unlock()
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
