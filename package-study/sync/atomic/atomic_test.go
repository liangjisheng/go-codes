package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic1(t *testing.T) {
	var a, b int32 = 13, 13
	var c int32 = 9
	res := atomic.CompareAndSwapInt32(&a, b, c)
	fmt.Println("swapped:", res) // true
	fmt.Println("替换的值:", c)      // 9
	fmt.Println("替换之后a的值:", a)   // 9

	fmt.Println()
	a, b = 13, 12
	old := atomic.SwapInt32(&a, b)
	fmt.Println("old的值:", old) // 13
	fmt.Println("替换之后a的值", a)  // 12

	fmt.Println()
	addValue := atomic.AddInt32(&a, 1)
	fmt.Println("a:", a)           // 13
	fmt.Println("增加之后:", addValue) // 13
	delValue := atomic.AddInt32(&a, -4)
	fmt.Println("a:", a)           // 9
	fmt.Println("减少之后:", delValue) // 9
}

func TestAtomic2(t *testing.T) {
	c1 := CommonCounter{} // 非并发安全
	dataCount(&c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	dataCount(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	dataCount(&c3)
}
