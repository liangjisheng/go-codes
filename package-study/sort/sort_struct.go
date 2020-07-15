package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int {
	return len(pw.people)
}

func (pw PersonWrapper) Swap(i, j int) {
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}

// 自定义Less方法,可以根据结构体的多个字段进行排序
func (pw PersonWrapper) Less(i, j int) bool {
	return pw.by(&pw.people[i], &pw.people[j])
}

func main() {
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
}
