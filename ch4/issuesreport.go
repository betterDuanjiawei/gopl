package main
//  go run issuesreport.go repo:golang/go is:open json decoder
import (
	"gopl.io/ch4/github"
	"text/template"
	"log"
	"os"
	"time"
)

// 表示当期值的标记,用点号(.)表示,
const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title|printf "%.64s"}}
Age: {{.CreatedAt| daysAgo}} days
{{end}}`
/**
const templ = `{{.TotalCount}} issues: // .表示 github.IssuesSearchResult, .TotalCount代表 TotalCount 代表 TotalCount 成员的值.
{{range .Items}}---------------------
Number: {{.Number}} // .表示 Items里面连续的元素
User: {{.User.Login}}
Title: {{.Title|printf "%.64s"}} // | 会把前一个操作当做下一个操作的输入,和 shell的管道类似. printf 在所有的模板中,就是内置函数 fmt.Sprintf 的同义词.
Age: {{.CreateAt| daysAgo}} days // daysAgo 使用 time.Since 将 CreatedAt转换为已过去的时间
{{end}}`
 */

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()/24)
}
var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo":daysAgo}).Parse(templ))

func main()  {
	// 通过模本输出结果需要两个步骤.1.需要解析模板并转换为内部的表示方法,然后在指定的输入上面执行.解析模板只需要执行一次. 创建并解析上面定义的文本模板 templ. 方法的链式调用:template.New 创建并返回一个新的模板,Funcs添加 daysAgo 到模板内部可以访问的函数列表中.然后返回这个模板对象,最后调用 Parse方法.
	// 模板通常是在编译器就固定下来的. template.Must 提供了一种便捷的错误处理方式.它接受一个模板和错误作为参数,检查错误是否为 nil.如果不是 nil,则宕机,然后返回这个模板.
	//report, err := template.New("report").Funcs(template.FuncMap{"daysAgo":daysAgo}).Parse(templ)
	//if err != nil {
	//	log.Fatal(err)
	//}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}