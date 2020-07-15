package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))            // unsafe.Pointer
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f)))) // *uint64
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// 通常情况下布尔和数字类型需要对齐到它们本身的大小(最多8个字节),
	// 其它的类型对齐到机器字大小.(64位的机器字大小为64位,8字节)
	fmt.Printf("%-30s%-30s%-30s%-50s\n",
		"Row", "Sizeof", "Alignof(对齐倍数)", "Offsetof(偏移量)")

	fmt.Printf("%-30s%-30d%-30d%-50s\n",
		"x", unsafe.Sizeof(x), unsafe.Alignof(x), "")
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.a", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.b", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.c", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
	fmt.Println()

	fmt.Printf("%#016x\n", Float64bits(1.0))
}
