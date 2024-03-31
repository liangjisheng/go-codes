package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// 当我们需要打印日志的时候经常会需要标示当前的代码位置信息，包括所在文件名，
// 行号，以及所在函数等等；特别是在处理log信息的时候
// go语言提供的runtime和reflect库可以帮助我们获取这些信息。下面是一个重写的
// log函数例子；自定义了一套log接口： ENTRY/EXIT/INFO/DEBUG等等，这些接口
// 都是相似，所以代码例子只给出了ENTRY和DEBUG

type myStruct struct {
}

func (m *myStruct) foo(p string) {
	ENTRY("")
	ENTRY("Param p=%s", p)
	DEBUG("Test %s %s", "Hello", "World")
}

// DEBUG ...
func DEBUG(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

// ENTRY ...
func ENTRY(formating string, args ...interface{}) {
	LOG("ENTRY", formating, args...)
}

// LOG ...
func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*myStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename)
	}

	log.Printf("%s:%d:%s: %s: %s\n", filename, line, funcname, level, fmt.Sprintf(formating, args...))
}

func main() {
	ss := myStruct{}
	ss.foo("hello")
}
