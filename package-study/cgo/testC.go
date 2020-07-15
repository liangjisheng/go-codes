package main

/*
#include <stdio.h>

void hello(const char *str)
{
    printf("%s\n", str);
}
*/
import "C" // 必须单起一行，且紧跟在注释行之后

func testC() {
	s := "Hello Cgo"
	cs := C.CString(s) // 字符串映射
	C.hello(cs)        // 调用c函数
	// defer C.free(unsafe.Pointer(cs)) // 释放内存
}
