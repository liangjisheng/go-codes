package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func regexp1() {
	// 直接使用字符串进行正则表达式匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Compile一个优化的Regexp结构体
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	// 查找第一个匹配字符串
	fmt.Println(r.FindString("peach punch"))
	// 返回第一个匹配字符串的索引
	fmt.Println(r.FindStringIndex("peach punch"))
	// 返回第一个完全匹配和局部匹配的字符串
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// 返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带All的函数返回所有的匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// 第二个参数来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println()

	// 可以使用Compile的变体MustCompile
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	// regexp包也可以用来替换部分字符串为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// 传递一个函数进去操作匹配字符串
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	fmt.Println()

	r, _ = regexp.Compile("log([\\d])+.([\\d])+.([\\d])+.([\\d])+\\+main_([\\d])+")
	fmt.Println(r.MatchString("log10.169.225.199+main_0"))
	r, _ = regexp.Compile("([\\d])+")
	fmt.Println(r.MatchString("1234"))
}
