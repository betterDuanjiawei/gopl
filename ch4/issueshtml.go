package main

//  go run issueshtml.go repo:golang/go is:open json decoder > issues.html
//	go run issueshtml.go repo:golang/go 3133 10535  > issues2.html

//#	State	User	Title
//3133	closed	ukai	html/template: escape xmldesc as &lt;?xml
//10535	open	dvyukov	x/net/html: void element <link> has child nodes
import (
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"os"
)
// file://Users/v_duanjiawei/go/src/gopl.io/ch4/issues.html
var issueList = template.Must(template.New("issueList").Parse(`
<h1>{{.TotalCount}} issue </h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range.Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}