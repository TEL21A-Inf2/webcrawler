package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/tel21a-inf2/webcrawler/htmlparser"
)

// Ein einfaches Programm, das den Benutzer nach einer URL fragt und
// die Links auf dieser Seite ausgibt.

func main() {
	fmt.Println("Bitte eine URL eingeben:")
	var rawUrl string
	fmt.Scanln(&rawUrl)

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	htmlDoc, err := htmlparser.FromUrl(rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Es wurden folgende Links gefunden:")
	htmlDoc.Links().
		Filter(func(link htmlparser.Hyperlink) bool { return link.IsValid() }).
		Filter(func(link htmlparser.Hyperlink) bool { return link.Url.Path != "" }).
		Each(func(link *htmlparser.Hyperlink) { link.Url = parsedUrl.ResolveReference(link.Url) }).
		Filter(func(link htmlparser.Hyperlink) bool { return link.Url.Hostname() == "de.wikipedia.org" }).
		PrintAll()
}
