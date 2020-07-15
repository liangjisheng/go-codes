package main

import (
	"fmt"
	"sort"
)

// 任何类型只要实现了sort.Interface接口，都可以使用sort中的方法排序
// 但要求元素索引为整数

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

// 实现sort.Interface
func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func Example() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

func main() {
	Example()
}
