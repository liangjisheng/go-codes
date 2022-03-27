package main

import (
	"fmt"
	"plugin"
)

func main() {
	ptr, err := plugin.Open("aplugin.so")
	if err != nil {
		fmt.Println(err)
	}

	Add, _ := ptr.Lookup("Add")
	sum := Add.(func(int, int) int)(5, 4)
	fmt.Println("Add", sum)

	Sub, _ := ptr.Lookup("Subtract")
	sub := Sub.(func(int, int) int)(9, 8)
	fmt.Println("Sub", sub)
}
