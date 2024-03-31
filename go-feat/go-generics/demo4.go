package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

//如何在Golang泛型中使用队列操作
//在现实开发过程中，我们有可能会需要一个队列去处理一些数据，在泛型中，我们可以抽取部分重复逻辑来实现

type queue[T any] []T

func (q *queue[T]) enqueue(v T) {
	*q = append(*q, v)
}

func (q *queue[T]) dequeue() (T, bool) {
	if len(*q) == 0 {
		var zero T
		return zero, false
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return r, true
}

func case11() {
	q := new(queue[int])
	q.enqueue(5)
	q.enqueue(6)
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}

//Golang泛型中新加入的一些约束包

type Vector[T constraints.Ordered] struct {
	x, y T
}

func (v *Vector[T]) Add(x, y T) {
	v.x += x
	v.y += y
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("{x: %v, y: %v}", v.x, v.y)
}

func NewVector[T constraints.Ordered](x, y T) *Vector[T] {
	return &Vector[T]{x: x, y: y}
}

func case12() {
	v := NewVector[float64](1, 2)
	v.Add(2, 3)
	fmt.Println(v)
}
