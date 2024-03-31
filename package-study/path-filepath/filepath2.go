package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func printFileAndDir(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

func findTxtDir(path string, info os.FileInfo, err error) error {
	ok, err := filepath.Match(`*.txt`, info.Name())
	if ok {
		fmt.Println(filepath.Dir(path), info.Name())
		// 遇到 txt 文件则继续处理所在目录的下一个目录
		// 注意会跳过子目录
		return filepath.SkipDir
	}
	return err
}

func filepath2() {
	err := filepath.Walk(".", printFileAndDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()

	err = filepath.Walk(".", findTxtDir)
	if err != nil {
		fmt.Println(err)
		return
	}
}
