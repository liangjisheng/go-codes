package main

import (
	"fmt"
	"unsafe"
)

// Go语言导出函数指针给c语言使用
// 还有一种使用方式, 就是传递函数指针, 因为GO函数无法取址, 因此需要写个中间函数做个转换操作

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lclibrary

#include "clibrary.h"

int callOnMeGo_cgo(int in); // 声明
*/
import "C"

//export callOnMeGo
func callOnMeGo(in int) int {
	return in + 1
}

func main() {
	fmt.Printf("Go.main(): calling C function with callback to us\n")
	// 使用unsafe.Pointer转换
	C.some_c_func((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))
}
