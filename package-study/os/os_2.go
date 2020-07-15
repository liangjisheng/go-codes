package main

import (
	"fmt"
	"os"
)

func main() {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}

		if m[key] != "" {
			return m[key]
		}
		return key
	}

	s := "hello,world"
	s1 := "$hello,$world $finish"
	// Expand用mapping 函数指定的规则替换字符串
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))

	s2 := "hello $GOROOT"
	fmt.Println(os.ExpandEnv(s2))
	fmt.Println(os.Getenv("GOROOT"))

	// 判断一个字符是否是路径分隔符
	fmt.Println(os.IsPathSeparator('/'))
	fmt.Println(os.IsPathSeparator('|'))
	fmt.Println()

	filemode, err := os.Stat("os_1.go")
	if err != nil {
		fmt.Println("os.Stat error", err)
		return
	} else {
		fmt.Println("filename:", filemode.Name())
		fmt.Println("filesize:", filemode.Size())
		fmt.Println("filemode:", filemode.Mode())
		fmt.Println("modtime:", filemode.ModTime())
		fmt.Println("IS_DIR:", filemode.IsDir())
		fmt.Println("SYS:", filemode.Sys())
	}
}
