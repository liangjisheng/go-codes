package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		/* 		if f.IsDir() {
			return nil
		} */
		println(path)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk returned %v\n", err)
	}
}

func main() {
	var dirPath string
	if os.Args[1] != "" {
		dirPath = os.Args[1]
	}

	getFileList(dirPath)
}
