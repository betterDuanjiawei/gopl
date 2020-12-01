package main

import (
	"html/template"
	"log"
	"os"
)

func main()  {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct{
		A string // 不受信任的纯文本 会转义
		B template.HTML // 受信任的 HTML 属于 html/template 包 不会转义
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}