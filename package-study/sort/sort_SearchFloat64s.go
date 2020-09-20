package main

import (
	"fmt"
	"sort"
)

func search2() {
	a := []string{"a", "c"}
	// 切片必须以升序排列，返回可以插入第二个参数的索引位置
	// 如果不存在，返回切片的长度
	i := sort.SearchStrings(a, "b")
	fmt.Println(i) // 1

	b := []string{"a", "b", "c", "d"}
	i = sort.SearchStrings(b, "b")
	fmt.Println(i) // 1

	// 为了精确查找，切片必须以升序方式进行排序
	// 切片c不是升序
	c := []string{"d", "c"}
	i = sort.SearchStrings(c, "b")
	fmt.Println(i) // 0

	// 切片d也不是升序
	d := []string{"c", "d", "b"}
	i = sort.SearchStrings(d, "b")
	fmt.Println(i) // 0
	fmt.Println()

	// func Reverse(data Interface) Interface
	// 实现对data的逆序排列
	data := []int{1, 2, 5, 3, 4}
	fmt.Println(data)
	// sort.Reverse(sort.IntSlice(data))
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println(data)
}
