package main

import (
	"embed"
	"fmt"
)

//go:embed test.txt hello.txt
//go:embed file/file.txt
var f embed.FS

//go:embed file
var d embed.FS

//go:embed file/*.txt
var pre embed.FS

func main() {
	// 嵌入为文件系统
	data, _ := f.ReadFile("test.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("hello.txt")
	fmt.Println(string(data))
	// 嵌入的时候文件是啥，这里要对应指定为相同的文件路径
	data, _ = f.ReadFile("file/file.txt")
	fmt.Println(string(data))

	data, _ = d.ReadFile("file/file.txt")
	fmt.Println(string(data))

	data, _ = pre.ReadFile("file/name.txt")
	fmt.Println(string(data))

	printDir("file")
}

func printDir(name string) {
	// 返回[]fs.DirEntry
	entries, err := d.ReadDir(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("dir:", name)
	for _, entry := range entries {
		// fs.DirEntry的Info接口会返回fs.FileInfo，这东西被从os移动到了io/fs，接口本身没有变化
		info, _ := entry.Info()
		fmt.Println("file name:", entry.Name(), "\tisDir:", entry.IsDir(), "\tsize:", info.Size())
	}
	fmt.Println()
}
