package main

import (
	"fmt"
	"sort"
)

// 满足 Interface 接口的类型可以被本包的函数进行排序。
// type Interface interface {
//     // Len 方法返回集合中的元素个数
//     Len() int
//     // Less 方法报告索引 i 的元素是否比索引 j 的元素小
//     Less(i, j int) bool
//     // Swap 方法交换索引 i 和 j 的两个元素的位置
//     Swap(i, j int)
// }

// IntSlice ...
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

func sort2() {
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
