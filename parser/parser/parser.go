package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func Parse(bodyStr io.Reader) ([]Recipe, error) {
	var doc *html.Node
	var err error

	if doc, err = html.Parse(bodyStr); err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil, nil
}
