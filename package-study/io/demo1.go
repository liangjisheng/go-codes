package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Ustr 自定义对象实现io包里面的Reader接口
type Ustr struct {
	s string // 数据流
	i int    // 读写位置
}

// NewUstr 根据字符串创建Ustr对象
func NewUstr(s string) *Ustr {
	return &Ustr{s, 0}
}

// 获取未读取部分的数据长度
func (s *Ustr) Len() int {
	return len(s.s) - s.i
}

// 实现Ustr类型的Read方法
func (s *Ustr) Read(p []byte) (n int, err error) {
	for ; s.i < len(s.s) && n < len(p); s.i++ {
		c := s.s[s.i]
		// 将小写字母转换为大写字母，然后写入p中
		if 'a' <= c && c <= 'z' {
			p[n] = c + 'A' - 'a'
		} else {
			p[n] = c
		}

		n++
	}

	if 0 == n {
		return n, io.EOF
	}
	return n, nil
}

// 然后就可以用 ReadFull 方法读取 Ustr 对象的数据了
func demo1() {
	s := NewUstr("Hello World!")
	buf := make([]byte, s.Len())

	n, err := io.ReadFull(s, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", string(buf))
	fmt.Println(n, err)
}

func demo2() {
	buf := bytes.NewBuffer([]byte("Hello World!"))
	b := make([]byte, buf.Len()) // Len()返回未读取数据的长度
	n, err := buf.Read(b)
	fmt.Printf("%s %v\n", b[:n], err)
	// Hello World!   <nil>

	buf.WriteString("ABCDEFG\n")
	buf.WriteTo(os.Stdout)
	// ABCDEFG

	n, err = buf.Read(b)
	fmt.Printf("%s %v\n", b[:n], err)
	// ABCDEFG
	// EOF

	n, err = buf.Write(b)
	fmt.Printf("%d %s %v\n", n, buf.String(), err)
	// 12   Hello World!   <nil>

	c, err := buf.ReadByte()
	fmt.Printf("%c %s %v\n", c, buf.String(), err)
	// H   ello World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c %s %v\n", c, buf.String(), err)
	// e   llo World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   bytes.Buffer: UnreadByte: previous operation was not a read
}

func demo3() {
	io.WriteString(os.Stdout, "Hello World!\n")
	// Hello World!

	r := strings.NewReader("Hello World!")
	b := make([]byte, 15)

	n, err := io.ReadAtLeast(r, b, 20)
	fmt.Printf("%q   %d   %v\n", b[:n], n, err)
	// ""   0   short buffer

	r.Seek(0, 0)
	b = make([]byte, 15)

	n, err = io.ReadFull(r, b)
	fmt.Printf("%q   %d   %v\n", b[:n], n, err)
	// "Hello World!"   12   unexpected EOF
}

func demo4() {
	r := strings.NewReader("Hello World!")
	buf := make([]byte, 32)

	n, err := io.CopyN(os.Stdout, r, 5) // Hello
	fmt.Printf("\n%d   %v\n\n", n, err) // 5   <nil>

	r.Seek(0, 0)
	n, err = io.Copy(os.Stdout, r)      // Hello World!
	fmt.Printf("\n%d   %v\n\n", n, err) // 12   <nil>

	r.Seek(0, 0)
	r2 := strings.NewReader("ABCDEFG")
	r3 := strings.NewReader("abcdefg")

	n, err = io.CopyBuffer(os.Stdout, r, buf) // Hello World!
	fmt.Printf("\n%d   %v\n", n, err)         // 12   <nil>
	fmt.Printf("buf:%v\n", buf)

	n, err = io.CopyBuffer(os.Stdout, r2, buf) // ABCDEFG
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>
	fmt.Printf("buf:%v\n", buf)

	n, err = io.CopyBuffer(os.Stdout, r3, buf) // abcdefg
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>
	fmt.Printf("buf:%v\n", buf)
}
