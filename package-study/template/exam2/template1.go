package main

import (
	"html/template"
	"os"
)

func template1() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	temp1 := template.New("template test")
	temp1, _ = temp1.Parse(`{{with $x := "output" | printf "%q"}}{{$x}}{{end}}`)
	temp1.Execute(os.Stdout, nil)
	os.Stdout.Write([]byte("\n"))

	temp2 := template.New("template test")
	temp2, _ = temp2.Parse(`{{with $x := "output"}}{{printf "%q" $x}}{{end}}`)
	temp2.Execute(os.Stdout, nil)
	os.Stdout.Write([]byte("\n"))

	temp3 := template.New("template test")
	temp3, _ = temp3.Parse(`{{with $x := "output"}}{{$x | printf "%q"}}{{end}}`)
	temp3.Execute(os.Stdout, nil)
	os.Stdout.Write([]byte("\n"))
}
