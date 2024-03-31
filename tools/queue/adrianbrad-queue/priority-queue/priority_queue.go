package main

import (
	"fmt"

	"github.com/adrianbrad/queue"
)

//Priority Queue is a data structure where the order of the elements is
//given by a comparator function provided at construction. Implemented
//using container/heap standard library package.

func main() {
	elems := []int{2, 3, 4}

	priorityQueue := queue.NewPriority(
		elems,
		func(elem, otherElem int) bool { return elem < otherElem },
	)

	containsTwo := priorityQueue.Contains(2)
	fmt.Println(containsTwo) // true

	size := priorityQueue.Size()
	fmt.Println(size) // 3

	empty := priorityQueue.IsEmpty()
	fmt.Println(empty) // false

	if err := priorityQueue.Offer(1); err != nil {
		// handle err
	}

	elem, err := priorityQueue.Get()
	if err != nil {
		// handle err
	}

	fmt.Printf("elem: %d\n", elem) // elem: 1
}
