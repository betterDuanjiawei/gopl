package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func main()  {

}

func soleTitle(doc *html.Node) (title string, err error)  {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// 没有宕机
		case bailout{} :
			// 预期的宕机
			err = fmt.Errorf("multiple title elements")
		default:
			// 未预期的宕机,继续宕机
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "",fmt.Errorf("no title element")
	}
	return title, nil
}