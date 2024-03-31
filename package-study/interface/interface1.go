package main

import "fmt"

// http://mp.weixin.qq.com/s?__biz=MjM5MjAwODM4MA==&mid=2650719488&idx=3&sn=bc59d9831f0c3b03de2285dc9a57b65a&chksm=bea6b4d389d13dc519890ebeb0bd27eeb0bad56acfce2ada77860eb38bf8acd95692bb4457c6&scene=0&xtrack=1#rd

// iface包含两个字段：tab 是接口表指针，指向类型信息；data 是数据指针，
// 则指向具体的数据。它们分别被称为动态类型和动态值。而接口值包括动态类型和动态值
// 接口值的零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为
// nil 的情况下，这个接口值就才会被认为 接口值 == nil

type coder interface {
	code()
}

type gopher struct {
	name string
}

func (g gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func main() {
	var c coder
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	var g *gopher
	fmt.Println(g == nil)

	c = g
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)
}
