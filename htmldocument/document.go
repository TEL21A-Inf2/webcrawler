package htmldocument

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/tel21a-inf2/webcrawler/htmlparser"
	"github.com/tel21a-inf2/webcrawler/httpreader"
	"golang.org/x/net/html"
)

// Repräsentiert ein HTML-Dokument.
// Enthält den Wurzelknoten des geparsten HMTL-Dokuments
// und eine Liste der Links aus diesem Dokument.
type HtmlDocument struct {
	rootNode *html.Node

	links []string
}

// Erzeugt ein neues Dokument aus einem io.Reader.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromReader(reader io.Reader) (*HtmlDocument, error) {
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return &HtmlDocument{root, nil}, nil
}

// Erzeugt ein neues Dokument aus einem String.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromString(source string) (*HtmlDocument, error) {
	return FromReader(strings.NewReader(source))
}

// Erzeugt ein neues Dokument aus einem String.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromBytes(data []byte) (*HtmlDocument, error) {
	return FromReader(bytes.NewReader(data))
}

// Erzeugt ein neues Dokument aus einer Datei.
// Die Links werden noch nicht initialisiert, sondern erst bei Bedarf.
func FromFile(path string) (*HtmlDocument, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return FromReader(bufio.NewReader(file))
}

// Erzeugt ein neues Dokument aus einer URL.
func FromUrl(url string) (*HtmlDocument, error) {
	htmldata, err := httpreader.Get(url)
	if err != nil {
		return nil, err
	}
	return FromBytes(htmldata)
}

// Liefert die Links, auf die das Dokument verweist.
// Erzeugt die Liste, falls sie noch nicht existiert.
func (doc *HtmlDocument) Links() []string {
	if doc.links == nil {
		doc.links = htmlparser.GetUrlList(doc.rootNode)
	}
	return doc.links
}
