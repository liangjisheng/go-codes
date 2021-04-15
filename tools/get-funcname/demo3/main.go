package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(runtime.FuncForPC(pc).Name())
	foo()
}

func foo() {
	fmt.Printf("我是 %s, %s 在调用我?\n",
		printMyName(), printCallerName())
	bar()
}

func bar() {
	fmt.Printf("我是 %s, %s 又在调用我?\n",
		printMyName(), printCallerName())
	// trace()
	trace2()
}

// 打印函数本身的名称
func printMyName() string {
	// 注意这里Caller的参数是1, 因为我们将业务代码封装成了一个函数
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// 可以打印调用者的名称
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// Caller可以返回函数调用栈的某一层的程序计数器、文件信息、行号。
// 0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推
// func Caller(skip int) (pc uintptr, file string, line int, ok bool)

// Callers用来返回调用站的程序计数器, 放到一个uintptr中
// 0 代表 Callers 本身，这和上面的Caller的参数的意义不一样，历史原因造成的。 1 才对应这上面的 0
// func Callers(skip int, pc []uintptr) int
func trace() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

// 上面的Callers只是或者栈的程序计数器，如果想获得整个栈的信息，可以使用
// CallersFrames函数，省去遍历调用FuncForPC
func trace2() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

// 在程序panic的时候，一般会自动把堆栈打出来，如果你想在程序中获取堆栈信息
// 可以通过debug.PrintStack()打印出来。比如你在程序中遇到一个Error,
// 但是不期望程序panic,只是想把堆栈信息打印出来以便跟踪调试
// 你可以使用debug.PrintStack()
func dumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// 利用堆栈信息还可以获取goroutine的id
func goroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
