package main

import (
	"fmt"
	"regexp"
)

func regexp6() {
	pat := `(abc)(def)(ghi)`
	reg := regexp.MustCompile(pat)
	// 获取正则表达式字符串
	fmt.Println(reg.String())
	// 获取分组数量
	fmt.Println(reg.NumSubexp())
	fmt.Println()

	// 获取分组名称
	pat = `(?P<Name1>abc)(def)(?P<Name3>ghi)`
	reg = regexp.MustCompile(pat)
	for i := 0; i <= reg.NumSubexp(); i++ {
		fmt.Printf("%d: %q\n", i, reg.SubexpNames()[i])
	}

	fmt.Println()
	// 获取字面前缀
	pat = `(abc1)(abc2)(abc3)`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) // abc1abc2abc3 true

	pat = `(abc1)|(abc2)|(abc3)`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) //  false

	pat = `abc1|abc2|abc3`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) // abc false
}
