package main

import (
	"log"
	"os"
)

// FileReader ...
type FileReader struct {
	f *os.File
}

// wire对provider的返回值个数和顺序有所规定:
// 1. 第一个参数是需要生成的依赖对象
// 2. 如果返回2个返回值，第二个参数必须是func()或者error
// 3. 如果返回3个返回值，第二个参数必须是func()，第三个参数则必须是error

// NewFileReader *FileReader 构造函数，第二个参数是cleanup function
func NewFileReader(filePath string) (*FileReader, func(), error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	fr := &FileReader{
		f: f,
	}
	fn := func() {
		log.Println("cleanup")
		fr.f.Close()
	}
	return fr, fn, nil
}

func main() {
	_, cleanup, err := InitializeFileReader("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	// do something
}
