package main

import (
	"fmt"
	"html/template"
	"os"
)

func template1() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}} Id is {{.ID}} Country is {{.Country}}")
	p := Person{UserName: "ljs", ID: 1, Country: "China"}
	fmt.Println(p)
	t.Execute(os.Stdout, p)
}

func template2() error {
	ljs := Person{UserName: "ljs", ID: 1, Country: "China"}
	fmt.Println(ljs)

	tmpl, err := template.ParseFiles("./tmp.html")
	if err != nil {
		fmt.Println("template.ParseFiles Error happened...")
		return err
	}

	err = tmpl.Execute(os.Stdout, ljs)
	if err != nil {
		fmt.Println("tmpl.Execute error")
		return err
	}
	return nil
}
