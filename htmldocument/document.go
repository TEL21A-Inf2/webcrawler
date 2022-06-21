package htmldocument

import (
	"strings"

	"github.com/tel21a-inf2/webcrawler/htmlparser"
	"golang.org/x/net/html"
)

// Repräsentiert ein HTML-Dokument.
// Enthält den Wurzelknoten des geparsten HMTL-Dokuments,
// den Quelltext und eine Liste der Links aus diesem Dokument.
type HtmlDocument struct {
	rootNode *html.Node

	source string
	links  []string
}

// Erzeugt ein neues Dokument aus einem String.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromString(source string) (*HtmlDocument, error) {
	root, err := html.Parse(strings.NewReader(source))
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{root, source, nil}, nil
}

// Liefert die Links, auf die das Dokument verweist.
// Erzeugt die Liste, falls sie noch nicht existiert.
func (doc *HtmlDocument) Links() []string {
	if doc.links == nil {
		doc.parseForLinks()
	}
	return doc.links
}

// Interne Hilfsfunktion: Parst die Links.
func (doc *HtmlDocument) parseForLinks() {
	links := htmlparser.GetUrlList(doc.rootNode)
	doc.links = links
}
