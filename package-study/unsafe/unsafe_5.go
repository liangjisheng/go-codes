package main

import (
	"fmt"
	"unsafe"
)

// 通过指针修改结构体字段
func main() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用(偏移量为0)
	up0 := uintptr(p)
	// 将通用指针转换为byte型指针
	pb := (*byte)(p)
	*pb = 10 // 赋值
	// 结构体内容改变
	fmt.Println(s)

	// 偏移到第2个字段
	up := up0 + unsafe.Offsetof(s.b)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 20
	fmt.Println(s)

	// 偏移到第3个字段
	up = up0 + unsafe.Offsetof(s.c)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 30
	fmt.Println(s)

	// 偏移到第4个字段
	up = up0 + unsafe.Offsetof(s.d)
	p = unsafe.Pointer(up)
	pi := (*int64)(p)
	*pi = 40
	fmt.Println(s)
}
