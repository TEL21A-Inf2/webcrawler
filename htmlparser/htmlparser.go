package htmlparser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func GetLinks(source string) []string {
	result := make([]string, 0)

	doc, err := html.Parse(strings.NewReader(source))

	if err != nil {
		log.Fatal(err)
	}

	// vgl. https://pkg.go.dev/golang.org/x/net/html#example-Parse
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					result = append(result, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return result
}
