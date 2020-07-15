package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	// repalce the @ by " at "
	return substrs[0] + " at " + substrs[1]
}

// 我们希望把@替换成at例如：astaxie at beego.me，如果要实现这样的功能，
// 我们就需要自定义函数来做这个功能。
// 每一个模板函数都有一个唯一值的名字，然后与一个Go函数关联

func template1() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
    {{range .Emails}}
	    an emails {{.|emailDeal}}
    {{end}}
    {{with .Friends}}
    {{range .}}
	    my friend name is {{.Fname}}
    {{end}}
    {{end}}
    `)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
