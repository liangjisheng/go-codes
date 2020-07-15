package main

// 满足 Interface 接口的类型可以被本包的函数进行排序。
// type Interface interface {
//     // Len 方法返回集合中的元素个数
//     Len() int
//     // Less 方法报告索引 i 的元素是否比索引 j 的元素小
//     Less(i, j int) bool
//     // Swap 方法交换索引 i 和 j 的两个元素的位置
//     Swap(i, j int)
// }

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

	c := sort.Reverse(a)          // 只是更改排序行为，并没有真正发生排序
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
	fmt.Println(d)                //&{0xc0000401d0}
}
