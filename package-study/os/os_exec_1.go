package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// func LookPath(file string) (string, error)
	// LookPath在环境变量中查找科执行二进制文件，如果file中包含一个斜杠，
	// 则直接根据绝对路径或者相对本目录的相对路径去查找
	f, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)
}
