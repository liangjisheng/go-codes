package main

import "fmt"

//go:generate go run main.go
//go:generate go version
func main() {
	fmt.Println("alice")
}

//执行下面的命令查看输出
//go generate -x
