package main

import (
	"bufio"
	"fmt"
	"strings"
)

func bufio1() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)

	fmt.Println(buf.Buffered()) // 0
	s, _ := buf.Peek(5)
	s[0], s[1], s[2] = 'a', 'b', 'c'
	fmt.Printf("%d   %q\n", buf.Buffered(), s) // 16   "abcDE"

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
}

func bufio2() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	buf := bufio.NewReaderSize(sr, 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFGHIJKLMNOP"   true   <nil>
	// "QRSTUVWXYZ"   false   <nil>
	// "1234567890"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部有一个换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF
}
