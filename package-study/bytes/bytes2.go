package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func byte6() {
	bs := [][]byte{
		[]byte("Hello World!"),
		[]byte("Hello 世界！"),
		[]byte("hello golang."),
	}

	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！."), r)
	}

	for _, b := range bs {
		// 删除b两边符合条件的字符
		fmt.Printf("%q\n", bytes.TrimFunc(b, f))
	}

	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimPrefix(b, []byte("Hello ")))
	}
}

func byte7() {
	b1 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	// 将b1包装成bytes.Reader对象
	// bytes.Reader 实现了如下接口：
	// io.ReadSeeker
	// io.ReaderAt
	// io.WriterTo
	// io.ByteScanner
	// io.RuneScanner
	rd := bytes.NewReader(b1)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)

	// Reset()将底层数据切换为b2，同时复位所有标记(读取位置等信息)
	rd.Reset(b2)
	rd.Read(buf)
	fmt.Printf("%q\n", buf)
	// Len()返回未读取部分的数据长度
	// Size()返回底层数据的总长度，返回值永远不变
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())
}

func byte8() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println()

	s1 := "hello, 世界"
	fmt.Println(len(s))                     // 返回字节长度13
	fmt.Println(utf8.RuneCountInString(s1)) // 返回码点长度，也就是字符长度9

	// utf8解码，得到每一个码点
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// 当range的对象是一个字符串时，隐式的调用了UTF-8解码，生成一个一个Unicode码点(rune)
	for i, r := range s1 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println()

	// []rune可以把字符串转换为Unicode码点
	fmt.Printf("% x\n", s1)
	r := []rune(s1)
	fmt.Printf("%x\n", r)

	// 如果将[]rune转换为一个字符串，会生成一个UTF8编码的字符串
	fmt.Println(string(r))

	// 将一个整数转为字符串时会把整数解释成rune值(rune的底层类型是int32)，生成UTF-8形式
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	// 如果rune不合法，会用错误字符替代显示
	fmt.Println(string(1234567))
}

func bytesJoin() {
	hello := "hello"
	helloBytes := []byte(hello)
	fmt.Println(helloBytes)

	world := "world"
	worldBytes := []byte(world)
	fmt.Println(worldBytes)

	helloWorld := [][]byte{helloBytes, worldBytes}
	fmt.Println(helloWorld)

	helloWorlds := bytes.Join(helloWorld, []byte{})
	fmt.Println(helloWorlds)
}
