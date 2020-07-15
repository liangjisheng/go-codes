package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	ID       int
	UserName string
	Country  string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ljs := Person{UserName: "ljs", ID: 1, Country: "China"}
	tmpl, err := template.ParseFiles("userall.tpl", "header.tpl", "center.tpl", "footer.tpl")
	if err != nil {
		fmt.Println("template.ParseFiles error")
		return
	}

	err = tmpl.Execute(w, ljs)
	if err != nil {
		fmt.Println("tmpl.Execute error")
		return
	}
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":9090", nil)
}
