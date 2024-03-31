package main

import (
	"flag"
	"fmt"
)

func main() {
	var id int
	var name string
	var male bool

	// flag.Parsed()返回是否已经解析命令行参数
	fmt.Println("parsed?=", flag.Parsed())

	// 设置flag参数 (变量指针，参数名，默认值，帮助信息)
	// 也可以用以下带返回值的方法代替，不过他们返回的是指针，比较麻烦点
	// Int(name string, value int, usage string) *int
	// String(name string, value string, usage string) *string
	// Bool(name string, value bool, usage string) *bool
	flag.IntVar(&id, "id", 123, "help msg for id")
	flag.StringVar(&name, "name", "default name", "help msg for name")
	flag.BoolVar(&male, "male", false, "help msg for male")

	flag.Parse() // 解析

	fmt.Println("parsed?=", flag.Parsed())

	// 获取非falg参数
	fmt.Println(flag.NArg()) // 输出非flag参数的个数
	fmt.Println("-------- Args start --------")
	// 返回解析之后剩下的非flag参数。（不包括命令名）.就是无法进行flag匹配的有哪些
	fmt.Println(flag.Args())
	for i, v := range flag.Args() {
		fmt.Printf("arg[%d] = (%s).\n", i, v)
	}
	fmt.Println("-------- Args end --------")

	// visit只包已经设置了的flag
	// 按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签
	fmt.Println("-------- visit flag start --------")
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)
	})
	fmt.Println("-------- visit flag end --------")

	// visitAll只包含所有的flag(包括未设置的)
	fmt.Println("-------- visitAll flag start --------")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)
	})
	fmt.Println("-------- visitAll flag end --------")
	fmt.Println()

	// flag参数
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("male = %t\n", male)
	fmt.Println()

	// flag参数默认值
	fmt.Println("-------- PrintDefaults start --------")
	flag.PrintDefaults()
	fmt.Println("-------- PrintDefaults end --------")
	fmt.Println()

	// 非flag参数个数
	fmt.Printf("NArg = %d\n", flag.NArg())
	// 已设置的flag参数个数
	fmt.Printf("NFlag = %d\n", flag.NFlag())
}
