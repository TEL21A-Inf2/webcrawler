package htmldocument

import (
	"strings"

	"github.com/tel21a-inf2/webcrawler/htmlparser"
	"golang.org/x/net/html"
)

type HtmlDocument struct {
	rootNode *html.Node

	source string
	links  []string
}

func FromString(source string) (*HtmlDocument, error) {
	root, err := html.Parse(strings.NewReader(source))
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{root, source, nil}, nil
}

func (doc *HtmlDocument) Links() []string {
	if doc.links == nil {
		doc.parseForLinks()
	}
	return doc.links
}

func (doc *HtmlDocument) parseForLinks() {
	links := htmlparser.GetUrlList(doc.rootNode)
	doc.links = links
}
