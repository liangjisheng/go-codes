package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

const (
	Separator     = os.PathSeparator     // 路径分隔符（分隔路径元素）
	ListSeparator = os.PathListSeparator // 路径列表分隔符（分隔多个路径）
)

// 下面两个函数主要用于将 Windows 路径分隔符转换为 Linux 路径分隔符，
// 处理完后再转换回去，只在 Windows 中有用，在 Linux 中没必要

// 将 path 中平台相关的路径分隔符转换为 '/'
// ToSlash(path string) string
// 将 path 中的 '/' 转换为系统相关的路径分隔符
// FromSlash(path string) string

func filepath1() {
	s := `http://www.site.com/a/b/c/d`
	u, _ := url.Parse(s)
	s = u.Path
	// 下面这句用于 Windows 系统
	s = filepath.FromSlash(s)
	fmt.Println(s) // /a/b/c/d 或 \a\b\c\d

	if err := os.MkdirAll(s[1:], 0755); err != nil {
		fmt.Println(err)
	}

	// 下面这句用于 Windows 系统
	s = filepath.ToSlash(s)
	fmt.Println(s) // /a/b/c/d
}
