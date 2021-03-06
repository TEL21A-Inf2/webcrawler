package htmlparser

import (
	"github.com/PuerkitoBio/goquery"
)

// Repräsentiert ein HTML-Dokument.
// Enthält den Wurzelknoten des geparsten HMTL-Dokuments
// und eine Liste der Links aus diesem Dokument.
type HtmlDocument struct {
	doc *goquery.Document

	links LinkList
	text  string
}

// Liefert die Links, auf die das Dokument verweist.
// Erzeugt die Liste, falls sie noch nicht existiert.
func (doc *HtmlDocument) Links() LinkList {
	if doc.links == nil {
		doc.links = GetLinks(doc.doc)
	}
	return doc.links
}

func (doc *HtmlDocument) Text() string {
	if doc.text == "" {
		doc.text = GetText(doc.doc)
	}
	return doc.text
}

func (doc *HtmlDocument) Contains(substring string) bool {
	return Contains(doc.doc, substring)
}
