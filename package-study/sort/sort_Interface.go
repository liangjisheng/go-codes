package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}

func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func main() {
	a := IntSlice{1, 3, 5, 7, 2}
	b := []float64{1.1, 2.3, 5.3, 3.4}
	c := []int{1, 3, 5, 4, 2}
	fmt.Println(sort.IsSorted(a)) // false
	if !sort.IsSorted(a) {
		sort.Sort(a)
	}

	if !sort.Float64sAreSorted(b) {
		sort.Float64s(b)
	}

	if !sort.IntsAreSorted(c) {
		sort.Ints(c)
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
