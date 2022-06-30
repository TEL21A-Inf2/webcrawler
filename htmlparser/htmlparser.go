package htmlparser

import (
	"github.com/PuerkitoBio/goquery"
)

// Liefert eine Liste der URLs, die im HTML-Baum von node als Links vorkommen.
func GetLinks(doc *goquery.Document) []Hyperlink {
	result := make([]Hyperlink, 0)

	addSelectionToResult := func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		result = append(result, Hyperlink{url, s.Text()})
	}

	doc.Find("a").
		Filter("a[href]").
		Each(addSelectionToResult)

	return result
}
