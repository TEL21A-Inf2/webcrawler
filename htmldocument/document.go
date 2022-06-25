package htmldocument

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tel21a-inf2/webcrawler/htmlparser"
	"github.com/tel21a-inf2/webcrawler/httpreader"
)

// Repräsentiert ein HTML-Dokument.
// Enthält den Wurzelknoten des geparsten HMTL-Dokuments
// und eine Liste der Links aus diesem Dokument.
type HtmlDocument struct {
	doc *goquery.Document

	links []string
}

// Erzeugt ein neues Dokument aus einem String.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromString(data string) (*HtmlDocument, error) {
	doc, err := httpreader.GoQueryDocumentFromString(data)
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{doc, nil}, nil
}

// Erzeugt ein neues Dokument aus einem String.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromBytes(data []byte) (*HtmlDocument, error) {
	doc, err := httpreader.GoQueryDocumentFromBytes(data)
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{doc, nil}, nil
}

// Erzeugt ein neues Dokument aus einer Datei.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromFile(path string) (*HtmlDocument, error) {
	doc, err := httpreader.GoQueryDocumentFromFile(path)
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{doc, nil}, nil
}

// Erzeugt ein neues Dokument aus einer URL.
func FromUrl(url string) (*HtmlDocument, error) {
	doc, err := httpreader.GoQueryDocumentFromUrl(url)
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{doc, nil}, nil
}

// Liefert die Links, auf die das Dokument verweist.
// Erzeugt die Liste, falls sie noch nicht existiert.
func (doc *HtmlDocument) Links() []string {
	if doc.links == nil {
		doc.links = htmlparser.GetUrlList(doc.doc)
	}
	return doc.links
}
