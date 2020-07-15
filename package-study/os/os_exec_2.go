package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func test_exec1() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// in all caps: "SOME INPUT"
	fmt.Printf("in all caps: %q\n", out.String())
}

func test_exec2() {
	//运行命令，并返回标准输出和标准错误
	// func (c *Cmd) CombinedOutput() ([]byte, error)
	// func (c *Cmd) Output() ([]byte, error)
	// 但他们两个不能同时使用
	cmd := exec.Command("ls")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

func test_exec3() {
	// StderrPipe返回一个pipe，这个管道连接到command的标准错误，
	// 当command命令退出时，Wait将关闭这些pipe
	// func (c *Cmd) StderrPipe() (io.ReadCloser, error)
	// StdinPipe返回一个连接到command标准输入的管道pipe
	// func (c *Cmd) StdinPipe() (io.WriteCloser, error)
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stdin.Write([]byte("some text"))
	if err != nil {
		fmt.Println(err)
		return
	}

	stdin.Close()
	cmd.Stdout = os.Stdout // 终端标准输出some text
	cmd.Start()
}

func test_exec4() {
	// StdoutPipe返回一个连接到command标准输出的管道pipe
	// func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
	cmd := exec.Command("ls")
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	test_exec1()
	fmt.Println()

	test_exec2()

	test_exec3()
	fmt.Println()

	test_exec4()
}
