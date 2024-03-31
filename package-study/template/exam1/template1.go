package main

import (
	"fmt"
	"html/template"
	"os"
)

func template1() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}} Id is {{.ID}} Country is {{.Country}}")
	p := Person{UserName: "alice", ID: 1, Country: "China"}
	fmt.Println(p)
	t.Execute(os.Stdout, p)
}

func template2() error {
	alice := Person{UserName: "alice", ID: 1, Country: "China"}
	fmt.Println(alice)

	tmpl, err := template.ParseFiles("./tmp.html")
	if err != nil {
		fmt.Println("template.ParseFiles Error happened...")
		return err
	}

	err = tmpl.Execute(os.Stdout, alice)
	if err != nil {
		fmt.Println("tmpl.Execute error")
		return err
	}
	return nil
}
