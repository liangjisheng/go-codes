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

// Person ...
type Person struct {
	Name string
	Age  int
}

// PersonWrapper ...
type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int      { return len(pw.people) }
func (pw PersonWrapper) Swap(i, j int) { pw.people[i], pw.people[j] = pw.people[j], pw.people[i] }

// 可以按 Age 排序
// func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

// 也可以自定义Less方法, 可以根据结构体的多个字段进行排序
func (pw PersonWrapper) Less(i, j int) bool { return pw.by(&pw.people[i], &pw.people[j]) }

func sort3() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}
	fmt.Println(people)

	// Age 递减排序
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age
	}})
	fmt.Println(people)

	// Name 递增排序
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return p.Name < q.Name
	}})
	fmt.Println(people)

	// 稳定排序
	sort.Stable(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age
	}})
}
