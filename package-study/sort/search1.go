package main

import (
	"fmt"
	"sort"
)

// GuessingGame ...
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func search() {
	a := []int{1, 2, 3, 4, 5}
	b := sort.Search(len(a), func(i int) bool { return a[i] >= 30 })
	fmt.Println(b) // 5, 查找不到， 返回a slice的长度，而不是-1
	c := sort.Search(len(a), func(i int) bool { return a[i] <= 3 })
	fmt.Println(c) //0，利用二分法进行查找，返回符合条件的最左边数值的index，即为０
	d := sort.Search(len(a), func(i int) bool { return a[i] == 3 })
	fmt.Println(d) // 2
	fmt.Println()

	GuessingGame()
}
