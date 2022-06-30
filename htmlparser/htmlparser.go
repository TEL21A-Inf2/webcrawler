package htmlparser

import (
	"github.com/PuerkitoBio/goquery"
)

// Liefert eine Liste der URLs, die im HTML-Baum von node als Links vorkommen.
func GetLinks(doc *goquery.Document) []Hyperlink {
	result := make([]Hyperlink, 0)

	doc.Find("a").Filter("a[href]").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		result = append(result, Hyperlink{url, s.Text()})
	})

	return result
}
