package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Ustr string

func (us Ustr) String() string {
	return strings.ToUpper(string(us))
}

func (us Ustr) GoString() string {
	return `"` + strings.ToUpper(string(us)) + `"`
}

func (u Ustr) Format(f fmt.State, c rune) {
	write := func(s string) {
		f.Write([]byte(s))
	}
	switch c {
	case 'm', 'M':
		write("旗标：[")
		for s := "+- 0#"; len(s) > 0; s = s[1:] {
			if f.Flag(int(s[0])) {
				write(s[:1])
			}
		}
		write("]")
		if v, ok := f.Width(); ok {
			write(" | 宽度：" + strconv.FormatInt(int64(v), 10))
		}
		if v, ok := f.Precision(); ok {
			write(" | 精度：" + strconv.FormatInt(int64(v), 10))
		}
	case 's', 'v': // 如果使用 Format 函数，则必须自己处理所有格式，包括 %#v
		if c == 'v' && f.Flag('#') {
			write(u.GoString())
		} else {
			write(u.String())
		}
	default: // 如果使用 Format 函数，则必须自己处理默认输出
		write("无效格式：" + string(c))
	}
}

func fmt1() {
	u := Ustr("Hello World!")
	// "-" 标记和 "0" 标记不能同时存在
	fmt.Printf("%-+ 0#8.5m\n", u) // 旗标：[+- #] | 宽度：8 | 精度：5
	fmt.Printf("%+ 0#8.5M\n", u)  // 旗标：[+ 0#] | 宽度：8 | 精度：5
	fmt.Println(u)                // HELLO WORLD!
	fmt.Printf("%s\n", u)         // HELLO WORLD!
	fmt.Printf("%#v\n", u)        // "HELLO WORLD!"
	fmt.Printf("%d\n", u)         // 无效格式：d
}
