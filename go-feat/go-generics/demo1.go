package main

import "fmt"

// M 声明一个map类型， 关键字是K 值是any --> any ==> interface{}
// 并且关节词不支持any，底层不支持
type M[K string, V any] map[K]V // 这里的K不支持any，由于底层map不支持，所以使用string

func demoMap() {
	m1 := M[string, int]{"key": 1}
	m1["key"] = 1

	m2 := M[string, string]{"key": "value"}
	m2["key"] = "new value"
	fmt.Println(m1, m2)
}

type C[T any] chan T

func demoChannel() {
	c1 := make(C[int], 10)
	c1 <- 1
	c1 <- 2

	c2 := make(C[string], 10)
	c2 <- "hello"
	c2 <- "world"

	fmt.Println(<-c1, <-c2)
}

// 定义一个any类型的参数
// T 就是any类型
//函数可以有一个额外的类型参数列表，它使用方括号，但看起来像一个普通的参数列表
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func demoSlice() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[float64]([]float64{1.01, 2.02, 3.03, 4.04, 5.05})
	printSlice[string]([]string{"alice", "lisi", "bob"})
	// 在编译器完全可以实现类型推导时，也可以省略显式类型
	printSlice([]int{66, 77, 88, 99, 100})
	printSlice([]float64{1.01, 2.02, 3.03, 4.04, 5.05})
	printSlice([]string{"alice", "lisi", "bob"})
}

//Addable 这个例子包含了一个类型约束。每个类型参数都有一个类型约束，就像每个普通参数都有一个类型：func F[T Constraint](p T) { ... }，
//类型约束是接口类型。该提案扩展了interface语法，新增了类型列表(type list)表达方式，专用于对类型参数进行约束
type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 | complex64 | complex128 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

//如果不用 interface 约束，直接使用的话，会得到如下的结果
//invalid operation: operator + not defined on a (variable of type T constrained by any)
//func add1[T any](a, b T) T {
//	return a + b
//}

func case1() {
	fmt.Println(add(1, 2))
	fmt.Println(add("hello", "world"))

	//fmt.Println(add1(1, 2))
}

//Addable2 在约束里，甚至可以放进去接口
type Addable2 interface {
	int | interface{}
}

func add2[T Addable2](a T) T {
	return a
}

func case2() {
	fmt.Println(add2(1))
}

//Addable3 去掉 string
//如果编译器通过类型推导得到的类型不在这个接口定义的类型约束列表中
//那个类型参数实例化将报错
type Addable3 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 | complex64 | complex128
}

func add3[T Addable3](a, b T) T {
	return a + b
}

func case3() {
	fmt.Println(add3(1, 2))
	//string does not implement Addable3
	//fmt.Println(add3("hello", "world"))
}

type MyType interface {
	int
}

//我们自己定义的带有类型列表的接口 MyType 将无法用作接口变量类型，如下代码将会报错
func case4() {
	//var n int = 6
	//var i MyType // interface contains type constraints
	//i = n
	//_ = i
}
