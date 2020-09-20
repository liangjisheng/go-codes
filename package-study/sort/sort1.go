package main

import (
	"fmt"
	"sort"
)

func sort1() {
	i := []int{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a := sort.IntSlice(i)
	fmt.Println(sort.IsSorted(a)) // false
	sort.Sort(a)
	fmt.Println(a)
	fmt.Println(sort.IsSorted(a)) // true

	b := sort.IntSlice{3}
	fmt.Println(sort.IsSorted(b)) // true

	c := sort.Reverse(a)          // 只是更改排序行为, 并没有真正发生排序
	fmt.Println(sort.IsSorted(c)) // false
	fmt.Println(c)
	sort.Sort(c)
	fmt.Println(c)
	fmt.Println(sort.IsSorted(c)) // true
	fmt.Println()

	d := sort.Reverse(c)
	fmt.Println(sort.IsSorted(d)) // false
	sort.Sort(d)
	fmt.Println(d)                // &{0xc0000401d0}
	fmt.Println(sort.IsSorted(d)) // true
	fmt.Println(d)                // &{0xc0000401d0}
}
