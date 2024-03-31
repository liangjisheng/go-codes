package main

import (
	"fmt"
	"regexp"
)

// func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
// func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

func re3() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$args')\n"), src, s)
	}
	fmt.Println(string(res))
}
