package main

import "fmt"

// 只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用
// recover，可以取到 panic 的错误信息，并且停止 panic 续发事件
// (Panicking Sequence),程序运行恢复正常.如果在延迟函数的外部
// 调用recover,就不能停止panic续发事件

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
