package main
// go run issueshtml.go repo:golang/go is:open json decoder > issues.html
import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
)
//  go run issues.go repo:golang/go is:open json decoder
func main()  {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}
}