package main

import (
	"fmt"

	"github.com/adrianbrad/queue"
)

//Blocking queue is a FIFO ordered data structure. Both blocking and non-blocking
//methods are implemented. Blocking methods wait for the queue to have available
//items when dequeuing, and wait for a slot to become available in case the queue
//is full when enqueuing. The non-blocking methods return an error if an element
//cannot be added or removed. Implemented using sync.Cond from the standard library.

func main() {
	elems := []int{1, 2}

	blockingQueue := queue.NewBlocking(elems, queue.WithCapacity(3))

	containsTwo := blockingQueue.Contains(2)
	fmt.Println(containsTwo) // true

	size := blockingQueue.Size()
	fmt.Println(size) // 2

	empty := blockingQueue.IsEmpty()
	fmt.Println(empty) // false

	if err := blockingQueue.Offer(1); err != nil {
		// handle err
	}

	elem, err := blockingQueue.Get()
	if err != nil {
		// handle err
	}

	fmt.Println("elem:", elem) // elem: 2

	blockingQueue.OfferWait(3)
	elem = blockingQueue.GetWait()
	fmt.Println("elem:", elem)
}
