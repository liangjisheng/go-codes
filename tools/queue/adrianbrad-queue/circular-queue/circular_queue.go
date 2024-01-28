package main

import (
	"fmt"

	"github.com/adrianbrad/queue"
)

//Circular Queue is a fixed size FIFO ordered data structure. When the queue is full,
//adding a new element to the queue overwrites the oldest element.

func main() {
	elems := []int{2, 3, 4}

	circularQueue := queue.NewCircular(elems, 3)

	containsTwo := circularQueue.Contains(2)
	fmt.Println(containsTwo) // true

	size := circularQueue.Size()
	fmt.Println(size) // 3

	empty := circularQueue.IsEmpty()
	fmt.Println(empty) // false

	if err := circularQueue.Offer(1); err != nil {
		// handle err
	}

	elem, err := circularQueue.Get()
	if err != nil {
		// handle err
	}

	fmt.Printf("elem: %d\n", elem) // elem: 1
}
