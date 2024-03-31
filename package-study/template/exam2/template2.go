package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func template2() {
	f1 := Friend{Fname: "Alice"}
	f2 := Friend{Fname: "Bob"}
	t := template.New("friendname example")
	t, _ = t.Parse(`hello {{.UserName}}!
    {{range .Emails}}
	    an email {{.}}
    {{end}}
    {{with .Friends}}
    {{range .}}
	    my friend name is {{.Fname}}
    {{end}}
    {{end}}
	`)
	p := Person{UserName: "alice",
		Emails:  []string{"Alice@beego.me", "Bob@beego.me"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
