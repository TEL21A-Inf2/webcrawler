package htmlparser

import (
	"github.com/PuerkitoBio/goquery"
)

// Liefert eine Liste der URLs, die im HTML-Baum von node als Links vorkommen.
func GetUrlList(doc *goquery.Document) []string {
	result := make([]string, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if exists {
			//			fmt.Printf("Link %d:\nURL: %s\nText: %s\n", i, url, text)
			result = append(result, url)
		}
	})

	return result
}
