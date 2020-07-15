package main

import (
	"bytes"
	"fmt"
)

func bytes1() {
	s1 := "Φφϕ kKK"
	s2 := "ϕΦφ KkK"

	// 看看 s1 里面是什么
	for _, c := range s1 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s2 里面是什么
	for _, c := range s2 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s1 和 s2 是否相似
	fmt.Println(bytes.EqualFold([]byte(s1), []byte(s2)))
}

func bytes2() {
	bs := [][]byte{
		[]byte("Hello World !"),
		[]byte("Hello 世界！"),
		[]byte("hello golang ."),
	}
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！. "), r)
	}
	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimFunc(b, f))
	}
	// "Hello World"
	// "Hello 世界"
	// "Hello Golang"
	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimPrefix(b, []byte("Hello ")))
	}
	// "World !"
	// "世界！"
	// "hello Golang ."
}

func bytes3() {
	b := []byte("  Hello   World !  ")
	fmt.Printf("%q\n", bytes.Split(b, []byte{' '}))
	// ["" "" "Hello" "" "" "World" "!" "" ""]
	fmt.Printf("%q\n", bytes.Fields(b))
	// ["Hello" "World" "!"]
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte(" !"), r)
	}
	fmt.Printf("%q\n", bytes.FieldsFunc(b, f))
	// ["Hello" "World"]
}

func bytes4() {
	b1 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	rd := bytes.NewReader(b1)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "World!"

	rd.Reset(b2)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())
	// Size:15, Len:9
}

func bytes5() {
	rd := bytes.NewBufferString("Hello World!")
	buf := make([]byte, 6)
	// 获取数据切片
	b := rd.Bytes()
	// 读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String()) // World!
	fmt.Printf("%s\n\n", b)         // Hello World!

	// 写入一部分数据，看看切片有没有变化
	rd.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", rd.String()) // World!abcdefg
	fmt.Printf("%s\n\n", b)         // Hello World!

	// 再读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String()) // abcdefg
	fmt.Printf("%s\n", b)           // Hello World!
}
