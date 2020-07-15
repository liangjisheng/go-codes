package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func SelfDefineSort() {
	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Sort(IntSlice(a))
	fmt.Println("After sorted:", a)
}

// sort.Stable() 稳定排序
type person struct {
	Name string
	Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	SelfDefineSort()

	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	// sort包分别定义了[]int和[]string的类型IntSlice和StringSlice
	// 并实现了各自的排序接口，sort.Ints()和sort.Strings()可以直接对[]int和[]string排序
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("After reversed:", a)

	myperson := personSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}

	sort.Stable(myperson)
	fmt.Println(myperson)
}
