package main

import (
	"fmt"

	"github.com/adrianbrad/queue"
)

//A linked queue, implemented as a singly linked list, offering O(1) time
//complexity for enqueue and dequeue operations. The queue maintains pointers
//to both the head (front) and tail (end) of the list for efficient operations
//without the need for traversal.

func main() {
	elems := []int{2, 3, 4}

	circularQueue := queue.NewLinked(elems)

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

	fmt.Printf("elem: %d\n", elem) // elem: 2
}
