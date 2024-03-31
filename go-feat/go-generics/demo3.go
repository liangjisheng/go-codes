package main

import "fmt"

//声明了一个可以存放任何类型的切片，叫做slice
type slice[T any] []T

//type any interface {
//	int | string
//}

func printSlice1[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

//和泛型函数一样，使用泛型类型时，首先要对其进行实例化，即显式为类型参数赋值类型。如果在类型定义时，
//将代码改成 vs:=slice{5,4,2,1} 那么你会得到如note1中的结果。因为编译器并没有办法进行类型推导，
//也就是表示它并不知道，你输出的是那种类型。哪怕你在interface里面定义了约束。哪怕你在接口中定义了类型约束
//type int, string，同样会报错，如note2所示
func case6() {
	// note1: cannot use generic type slice[T interface{}] without instantiation
	// note2: cannot use generic type slice[T any] without instantiation
	vs := slice[int]{5, 4, 2, 1}
	//vs := slice{5, 4, 2, 1}
	printSlice1(vs)

	vs2 := slice[string]{"hello", "world"}
	//vs2 := slice{"hello", "world"}
	printSlice1(vs2)
}

//利用泛型实现最大值最小值函数
type minmax interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func max[T minmax](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T minmax](a []T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}

func case7() {
	vi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := max(vi)
	fmt.Println(result)

	vi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = min(vi)
	fmt.Println(result)
}

//使用Golang泛型自带的comparable约束
//使用 any 作为类型约束的时候会报如下错误
//invalid operation: e == v (type parameter T is not comparable with ==)
//因为不是所有的类型都可以==比较
//所以Golang内置提供了一个comparable约束，表示可比较的

//func findFunc[T any](a []T, v T) int {
//	for i, e := range a {
//		if e == v {
//			return i
//		}
//	}
//	return -1
//}

func findFunc1[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

func case8() {
	//fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(findFunc1([]int{1, 2, 3, 4, 5, 6}, 5))
}

//在泛型中操作指针

func pointerOf[T any](v T) *T {
	return &v
}

func case9() {
	sp := pointerOf("foo")
	fmt.Println(*sp)

	ip := pointerOf(123)
	fmt.Println(*ip)
	*ip = 234
	fmt.Println(*ip)
}

//Golang泛型中如何操作map
//在现实开发过程中，我们往往需要对slice中数据的每个值进行单独的处理，比如说需要对其中数值转换为平方值
//在泛型中，我们可以抽取部分重复逻辑作为map函数

func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func case10() {
	vi := []int{1, 2, 3, 4, 5, 6}
	vs := mapFunc(vi, func(v int) string {
		return "<" + fmt.Sprint(v*v) + ">"
	})
	fmt.Println(vs)
}
