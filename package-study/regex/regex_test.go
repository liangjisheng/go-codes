package main

import (
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	// 查找以@开头,以空格结尾,中间不包含空格的字符串
	reg := regexp.MustCompile(`@[^ ]* `)
	str := "@xxx xxx "
	t.Log(reg.FindAllString(str, -1))
}

func TestRegex1(t *testing.T) {
	reg := regexp.MustCompile(`[\w-]+:`)
	str := "GOPATH:"
	t.Log(reg.FindAllString(str, -1))
}

func TestRegex2(t *testing.T) {
	reg := regexp.MustCompile(`(\S)+`)
	str := "rm rm"
	res := reg.FindAllString(str, -1)

	for _, v := range res {
		t.Log(v)
	}
}

func TestRegex3(t *testing.T) {
	str := "xxx"

	//reg := regexp.MustCompile(`^([0-9a-z]+[\.\-_])*[0-9a-z]+$`)
	reg := regexp.MustCompile(`^([0-9a-z]+_)*[0-9a-z]+$`)
	res := reg.FindAllString(str, -1)

	regLen := regexp.MustCompile(`^[0-9a-z_]{3,20}$`)
	resLen := regLen.FindAllString(str, -1)

	t.Log("res:", res)
	for _, v := range res {
		t.Log(v)
	}

	t.Log("resLen:", resLen)
	for _, v := range resLen {
		t.Log(v)
	}
}
