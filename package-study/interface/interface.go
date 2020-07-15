package main

import "fmt"

// http://mp.weixin.qq.com/s?__biz=MjM5MjAwODM4MA==&mid=2650719488&idx=3&sn=bc59d9831f0c3b03de2285dc9a57b65a&chksm=bea6b4d389d13dc519890ebeb0bd27eeb0bad56acfce2ada77860eb38bf8acd95692bb4457c6&scene=0&xtrack=1#rd

type coder interface {
	code()
	debug()
}

type gopher struct {
	language string
}

func (p gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	x := 200
	var any interface{} = x
	fmt.Println(any)

	g := gopher{"Go"}
	var c coder = g
	fmt.Println(c)
}
