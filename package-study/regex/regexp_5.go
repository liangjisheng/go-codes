package main

import (
	"fmt"
	"regexp"
)

func regexp5() {
	pat := `(((abc.)def.)ghi)`
	src := `abc-def-ghi abc+def+ghi`
	fmt.Println(regexp.MatchString(pat, src))

	// 将pat中的正则表达式元字符转义成普通字符
	fmt.Println(regexp.QuoteMeta(pat))

	// 第一匹配和最长匹配
	b := []byte("abc1def1")
	pat = `abc1|abc1def1`
	reg1 := regexp.MustCompile(pat)      // 第一匹配
	reg2 := regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))     // abc1
	fmt.Printf("%s\n", reg2.Find(b))     // abc1def1

	b = []byte("abc1def1")
	pat = `(abc|abcdef)*1`
	reg1 = regexp.MustCompile(pat)      // 第一个匹配
	reg2 = regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))    // abc1
	fmt.Printf("%s\n", reg2.Find(b))    // abc1
}
