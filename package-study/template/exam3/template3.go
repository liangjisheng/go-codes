package main

import (
	"fmt"
	"html/template"
	"os"
)

func template3() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()

	// 我们执行s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，
	// 所以不会输出任何的东西
	s1.Execute(os.Stdout, nil)
}
