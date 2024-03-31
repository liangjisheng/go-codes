package filedir

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"syscall"
)

//一次性读取全部内容

func ReadFile1(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func ReadFile2(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func ReadFile3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func ReadFile4(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

//每次只读取一行

func ReadFile5(filename string) {
	// 创建句柄
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)

	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

func ReadFile6(filename string) {
	// 创建句柄
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)

	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

//每次仅读取一行数据，可以解决内存占用过大的问题，但要注意的是，并不是所有的文件都有换行符 \n
//因此对于一些不换行的大文件来说，还得再想想其他办法
//每次只读取固定字节数

func ReadFile7(filename string) {
	// 创建句柄
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)

	// 每次读取 1024 个字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

//使用 syscall 库

func ReadFile8(filename string) {
	fd, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Failed on open: ", err)
	}
	defer syscall.Close(fd)

	var wg sync.WaitGroup
	wg.Add(2)
	dataChan := make(chan []byte)
	go func() {
		wg.Done()
		for {
			data := make([]byte, 100)
			n, _ := syscall.Read(fd, data)
			if n == 0 {
				break
			}
			dataChan <- data
		}
		close(dataChan)
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-dataChan:
				if !ok {
					return
				}

				fmt.Printf(string(data))
			default:

			}
		}
	}()
	wg.Wait()
}
