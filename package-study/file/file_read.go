package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "test.txt"
	Ioutil(filename)
	fmt.Println()
	OsIoutil(filename)
	fmt.Println()
	FileOpen(filename)
	fmt.Println()
	BufioRead(filename)
}

func Ioutil(name string) {
	if content, err := ioutil.ReadFile(name); err == nil {
		result := strings.Replace(string(content), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:", result)
	}
}

func OsIoutil(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		if content, err := ioutil.ReadAll(fileObj); err == nil {
			result := strings.Replace(string(content), "\n", "", 1)
			fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file:", result)
		}
	}
}

func FileOpen(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		// 在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
		buf := make([]byte, 1024)
		if n, err := fileObj.Read(buf); err == nil {
			fmt.Println("The number of bytes read:"+strconv.Itoa(n), "Buf length:"+strconv.Itoa(len(buf)))
			result := strings.Replace(string(buf[:50]), "\n", "", 1)
			fmt.Println("Use os.Open and File's Read method to read a file:", result)
		}
	}
}

func BufioRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		// 一个文件对象本身是实现了io.Reader的 使用bufio.NewReader
		// 去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
		reader := bufio.NewReader(fileObj)
		// 使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
		if result, err := reader.ReadString(byte('@')); err == nil {
			fmt.Println("使用ReadSlince相关方法读取内容:", result)
		}
		// 注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
		// 需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
		buf := make([]byte, 1024)
		// 读取Reader对象中的内容到[]byte类型的buf中
		if n, err := reader.Read(buf); err == nil {
			fmt.Println("The number of bytes read:" + strconv.Itoa(n))
			//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
			fmt.Println("Use bufio.NewReader and os.Open read file contents to a []byte:", string(buf))
		}
	}
}
