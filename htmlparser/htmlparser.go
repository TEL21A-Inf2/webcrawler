package htmlparser

import (
	"github.com/PuerkitoBio/goquery"
)

// Repräsentiert einen Hyperlink bestehend aus der Url und dem Link-Text.
type Hyperlink struct {
	url, text string
}

// Liefert eine Liste der URLs, die im HTML-Baum von node als Links vorkommen.
func GetLinks(doc *goquery.Document) []Hyperlink {
	result := make([]Hyperlink, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if exists {
			result = append(result, Hyperlink{url, s.Text()})
		}
	})

	return result
}
