package main

import (
	"fmt"
	"runtime"
)

func main() {
	hello()
	fmt.Println("main func:", CurrentFuncName())
}

func hello() {
	fmt.Println("hello func:", CurrentFuncName())
}

// CurrentFuncName 获取当前函数名
func CurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
