package main

import (
	"fmt"
	"regexp"
)

func regexp8() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)
	src := []byte(`abc-def-ghi abc+def+ghi`)
	template := []byte(`$0   $1   $2   $3`)

	// 替换第一次匹配结果
	match := reg.FindSubmatchIndex(src)
	fmt.Printf("%v\n", match)
	dst := reg.Expand(nil, template, src, match)
	fmt.Printf("%s\n\n", dst)

	// 替换所有匹配结果
	for _, match := range reg.FindAllSubmatchIndex(src, -1) {
		fmt.Printf("%v\n", match)
		dst := reg.Expand(nil, template, src, match)
		fmt.Printf("%s\n", dst)
	}
}
