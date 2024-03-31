package main

import "C"
import "github.com/rainycape/dl"

func main() {
	lib, err := dl.Open("./libhello.so", 0)
	if err != nil {
		panic(err)
	}
	defer lib.Close()
	var SayHello func(src *C.char) // 定义函数变量匹配 libhello 中的 SayHello 函数
	lib.Sym("SayHello", &SayHello) // 定位 SayHello 函数地址
	SayHello(C.CString("hello world"))
}
