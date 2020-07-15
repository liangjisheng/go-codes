package main

import "fmt"

func forMap() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	// go 始终利用同一块内存来接收集合中的一个值，只是在每次循环的时候重新赋值而已
	for index, value := range slice {
		fmt.Printf("%p\n", &value)
		myMap[index] = &value
	}
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}

	fmt.Println()
	for index, value := range slice {
		num := value
		myMap[index] = &num
	}
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
