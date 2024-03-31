package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//This looks like reasonable code, but under certain circumstances
//it fails in a surprising way:

//The race detector shows the problem: an unsynchronized read and write
//of the variable t from different goroutines. If the initial timer
//duration is very small, the timer function may fire before the main
//goroutine has assigned a value to t and so the call to t.Reset is made
//with a nil t.

//To fix the race condition we change the code to read and write the
//variable t only from the main goroutine:

func demo1() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})

	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

//Here the main goroutine is wholly responsible for setting and resetting
//the Timer t and a new reset channel communicates the need to reset the
//timer in a thread-safe way.
