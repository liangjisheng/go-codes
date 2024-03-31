package main

import (
	"fmt"
	"log"
	"os"
)

// os.Truncate("test.txt", 100)
// 裁剪一个文件到100个字节
// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
// 如果文件本来超过100个字节，则超过的字节会被抛弃。
// 这样我们总是得到精确的100个字节的文件。
// 传入0则会清空文件

func main() {
	filename := "test.txt"

	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newFile)
	newFile.Close()

	// 如果文件不存在则返回error
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("FileName:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permission:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("System interface type: %T\n", fileInfo.Sys())
	fmt.Println("System info: %+v\n\n", fileInfo.Sys())
}
