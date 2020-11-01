package main

// Go语言调用C函数例子

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void hello(const char *str)
{
    printf("%s\n", str);
}
*/
import "C" // 必须单起一行，且紧跟在注释行之后

import (
	"fmt"
	"unsafe"
)

func testC() {
	s := "Hello Cgo"
	// 使用C.CString创建的字符串需要手动释放
	cs := C.CString(s) // 字符串映射
	C.hello(cs)        // 调用c函数
	C.free(unsafe.Pointer(cs))
	fmt.Println("call C.sleep for 3s")
	C.sleep(3)
}

func main() {
	testC()
}
