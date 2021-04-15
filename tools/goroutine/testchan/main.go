package main

import (
	"fmt"
	"time"
)

func sendValue(data chan int) {
	i := new(int)
	for {
		time.Sleep(time.Second)
		data <- *i
		*i++
	}
}

func recvValue(data chan int) {
	for i := range data {
		fmt.Println("recvValue:", i)
	}
}

func sendPoint(data chan *int) {
	i := new(int)
	for {
		time.Sleep(time.Second)
		data <- i
		*i++
	}
}

func recvPoint(data chan *int) {
	for i := range data {
		fmt.Println("recvPoint:", *i)
	}
}

func main() {
	dataValue := make(chan int, 10)
	go sendValue(dataValue)
	go recvValue(dataValue)
	time.Sleep(time.Second * 3)
	close(dataValue)

	// dataPoint := make(chan *int, 10)
	// go sendPoint(dataPoint)
	// go recvPoint(dataPoint)
	// time.Sleep(time.Second * 3)
	// close(dataPoint)
}
