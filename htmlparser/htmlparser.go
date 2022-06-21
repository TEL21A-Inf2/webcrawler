package htmlparser

import (
	"errors"

	"golang.org/x/net/html"
)

// Prüft, ob der gegebenen Knoten ein Anchor-Knoten ist.
func isAnchorNode(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

// Liefert den Wert eines href-Attributs aus dem Knoten, falls eines existiert.
// Liefert einen error, falls kein href-Attribut existiert.
func getHref(node *html.Node) (string, error) {
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			return attribute.Val, nil
		}
	}
	return "", errors.New("node has no href attribute")
}

func GetUrlList(node *html.Node) []string {
	result := make([]string, 0)

	// Versuchen, aus node eine URL zu ziehen.
	if isAnchorNode(node) {
		url, err := getHref(node)
		if err == nil {
			result = append(result, url)
		}
	}

	// Rekursiv Links aus allen Kindern sammeln.
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		result = append(result, GetUrlList(childNode)...)
	}
	return result
}
