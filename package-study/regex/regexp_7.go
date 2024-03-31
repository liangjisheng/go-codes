package main

import (
	"fmt"
	"regexp"
)

func regexp7() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)
	src := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找第一个匹配结果
	// 返回第一个匹配到的结果(结果以切片形式返回)
	fmt.Printf("%s\n", reg.Find(src))
	fmt.Println()

	// 查找第一个匹配结果及其分组字符串
	// 返回第一个匹配到的结果及其分组内容(结果以切片形式返回)
	// 返回值中的第 0 个元素是整个正则表达式的匹配结果，后续元素是各个分组的
	// 匹配内容，分组顺序按照“(”的出现次序而定
	first := reg.FindSubmatch(src)
	for i := 0; i < len(first); i++ {
		fmt.Printf("%d: %s\n", i, first[i])
	}
}
