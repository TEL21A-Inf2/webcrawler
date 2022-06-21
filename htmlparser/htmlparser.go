package htmlparser

import (
	"log"

	"github.com/tel21a-inf2/webcrawler/htmldocument"
)

// Liefert alle Urls, die in source als Links vorkommen.
// source muss dafür gültiger HTML-Code sein.
func GetLinks(source string) []string {
	doc, err := htmldocument.FromString(source)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Links()
}
