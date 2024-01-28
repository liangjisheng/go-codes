package main

import (
	"fmt"
	"slices"
)

//Clip：从切片中删除未使用的容量，返回 s[:len(s):len(s)]
//Clone：拷贝切片的副本，切片元素是使用赋值复制的，是浅拷贝
//Grow：增加切片的容量，至少增加 n 个元素的空间。如果 n 是负数或者太大，无法分配内存，就会导致产生 panic
//Replace：用所传入的参数替换对应的元素，并返回修改后的切片
//IsSorted：检查所传入的切片是否以升序排序
//IsSortedFunc：同上，可传自定义函数
//Sort：按升序对任意有序类型的切片进行排序
//SortFunc：同上，可传自定义函数
//SortStableFunc：对所传入的切片进行排序，同时保持相等元素的原始顺序，使用较少的元素进行比较

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.IsSorted(s1))
	s2 := slices.Clone(s1)
	fmt.Println(s2)

	slices.Reverse(s1)
	fmt.Println(s1)

	slices.Sort(s1)
	fmt.Println(s1)

	maxNum := slices.Max(s1)
	fmt.Println(maxNum)
	minNum := slices.Min(s1)
	fmt.Println(minNum)

	s3 := slices.Delete(s1, 0, 1)
	fmt.Println(s3)

	idx := slices.Index(s3, 4)
	fmt.Println(idx)

	contain := slices.Contains(s3, 4)
	fmt.Println(contain)

	//BinarySearch：在已排序的切片中搜索目标，并返回找到目标的位置，或者目标在排序顺序中出现的位置；
	//函数会返回一个 bool 值，表示是否真的在切片中找到目标。切片必须按递增顺序排序。
	//BinarySearchFunc：同上类似用法，区别在于可以传自己定义的比较函数
	n, found := slices.BinarySearch(s3, 5)
	fmt.Printf("n %d, found %v\n", n, found)

	//Compact：将连续运行的相等元素替换为单个副本。类似于 Unix 的 uniq 命令。
	//该函数会直接修改切片的元素，它不会创建新切片。
	//CompactFunc：同上类似用法，区别在于可传自定义函数进行比较
	s1 = []int{1, 1, 2, 2, 3, 4, 5}
	s2 = slices.Compact(s1)
	fmt.Println(s2)

	ss1 := []string{"alice", "bob"}
	ss2 := []string{"Alice", "Bob"}
	equal := slices.Equal(ss1, ss2)
	fmt.Println("equal", equal)
}
