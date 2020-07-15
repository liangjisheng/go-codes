package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// read1()-read4()读取速度依次变慢

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return string(fd)
}

func read2(path string) string {
	fi, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(fi)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	f := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	chunks := make([]byte, 1024, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	return string(chunks)
}

func read4(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buf := make([]byte, 1024)
	chunks := make([]byte, 1024, 1024)

	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	return string(chunks)
}

func main() {
	logPath := "./test.log"

	start := time.Now()
	read1(logPath)
	t1 := time.Now()
	fmt.Printf("read1(): Cost time %v\n", t1.Sub(start).Nanoseconds())

	read2(logPath)
	t2 := time.Now()
	fmt.Printf("read2(): Cost time %v\n", t2.Sub(t1).Nanoseconds())

	read3(logPath)
	t3 := time.Now()
	fmt.Printf("read3(): Cost time %v\n", t3.Sub(t2).Nanoseconds())

	read4(logPath)
	t4 := time.Now()
	fmt.Printf("read4(): Cost time %v\n", t4.Sub(t3).Nanoseconds())
}
