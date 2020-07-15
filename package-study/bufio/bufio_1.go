package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)

	fmt.Println(buf.Buffered())
	s, _ := buf.Peek(5)
	s[0], s[1], s[2] = 'a', 'b', 'c'
	fmt.Printf("%d %q\n", buf.Buffered(), s)

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d %q %v\n", buf.Buffered(), b[:n], err)
	}
}
