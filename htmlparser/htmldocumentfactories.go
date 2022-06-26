package htmlparser

import (
	"github.com/tel21a-inf2/webcrawler/httpreader"
)

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
