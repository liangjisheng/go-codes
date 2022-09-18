package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func exec1() {
	command := flag.String("cmd", "pwd", "Set the command")
	args := flag.String("args", "", "Set the args.(separated by space)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[-cmd <command>][-args <the arguments>(separated by space)]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Println("Command:", *command)
	fmt.Println("Arguments:", *args)
	var argArray []string
	if *args != "" {
		argArray = strings.Split(*args, " ")
	} else {
		argArray = make([]string, 0)
	}

	cmd := exec.Command(*command, argArray...)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s(Command: %s, Arguments: %s)", err, *command, *args)
		return
	}

	fmt.Fprintf(os.Stdout, "Result: %s", buf)
}

func exec2() {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}

		if m[key] != "" {
			return m[key]
		}
		return key
	}

	s := "hello,world"
	s1 := "$hello,$world $finish"
	// Expand用mapping 函数指定的规则替换字符串
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))

	s2 := "hello $GOROOT"
	fmt.Println(os.ExpandEnv(s2))
	fmt.Println(os.Getenv("GOROOT"))

	// 判断一个字符是否是路径分隔符
	fmt.Println(os.IsPathSeparator('/'))
	fmt.Println(os.IsPathSeparator('|'))
	fmt.Println()

	filemode, err := os.Stat("os_1.go")
	if err != nil {
		fmt.Println("os.Stat error", err)
		return
	} else {
		fmt.Println("filename:", filemode.Name())
		fmt.Println("filesize:", filemode.Size())
		fmt.Println("filemode:", filemode.Mode())
		fmt.Println("modtime:", filemode.ModTime())
		fmt.Println("IS_DIR:", filemode.IsDir())
		fmt.Println("SYS:", filemode.Sys())
	}
}

func exec3() {
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

func exec4() {
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

func exec5() {
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

func exec6() {
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

func shell(str string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", str)
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(res), nil
}
