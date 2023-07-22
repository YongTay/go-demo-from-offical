package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	// 手动处理panic
	t1, err := t1.Parse("Value is {{.}}\n")
	if nil != err {
		panic(err)
	}
	// Must函数会帮忙处理panic
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(t).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Joe",
	})
	t2.Execute(os.Stdout, struct {
		Name string
	}{
		Name: "Jane",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, false)

	t4 := Create("t4", "Range: {{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
}
