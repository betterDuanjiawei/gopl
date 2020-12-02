package main
// cannot find package "golang.org/x/net/html" in any of:  https://blog.csdn.net/u014374009/article/details/105964286/
import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main()  {
	doc, err := html.Parse(os.Stdin)
	fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	}
	return links
}