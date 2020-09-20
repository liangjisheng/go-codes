package main

import (
	"fmt"
	"sort"
)

// Reverse ...
type Reverse struct {
	sort.Interface
}

// Less 只是将其中的 Interface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func sortReverse() {
	ints := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(ints)
	fmt.Println("After sort by Ints:", ints)

	doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.4, 1.8}
	sort.Float64s(doubles)
	fmt.Println("after sort by Float64s:", doubles)

	strings := []string{"hello", "good", "students", "morning", "people", "world"}
	sort.Strings(strings)
	fmt.Println("after sort by Strings:", strings)

	ipos := sort.SearchInts(ints, 5)
	fmt.Printf("pos of 5 is %d th\n", ipos)
	dpos := sort.SearchFloat64s(doubles, 20.1)
	fmt.Printf("pos of 5.0 is %d th\n", dpos)
	fmt.Printf("doubles is asc? %v\n", sort.Float64sAreSorted(doubles))

	(sort.Float64Slice(doubles)).Sort()
	fmt.Println("after sort by Sort:", doubles)

	sort.Sort(Reverse{sort.Float64Slice(doubles)})
	fmt.Println("after sort by Reversed Sort:", doubles)
}
