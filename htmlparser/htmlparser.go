package htmlparser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Liefert eine Liste der URLs, die im HTML-Baum von node als Links vorkommen.
func GetLinks(doc *goquery.Document) LinkList {
	result := make(LinkList, 0)

	addSelectionToResult := func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		result = append(result, NewHyperlink(url, s.Text()))
	}

	doc.Find("a").
		Filter("a[href]").
		Each(addSelectionToResult)

	return result
}

func GetText(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Text())
}

func Contains(doc *goquery.Document, substring string) bool {
	return strings.Contains(GetText(doc), substring)
}
