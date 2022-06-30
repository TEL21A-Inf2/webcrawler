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
		Filter(func(link htmlparser.Hyperlink) bool { return link.IsValid() }).                            // Nur valide Links
		Filter(func(link htmlparser.Hyperlink) bool { return link.Url.Path != "" }).                       // Keine Links ohne Pfadangabe (z.B. keine Links innerhalb der selben Webseite)
		Each(func(link *htmlparser.Hyperlink) { link.Url = parsedUrl.ResolveReference(link.Url) }).        // Alle Links absolut machen
		Filter(func(link htmlparser.Hyperlink) bool { return link.Url.Hostname() == "de.wikipedia.org" }). // Nur Links innerhalb der deutschen Wikipedia-Seite
		PrintAll()                                                                                         // Alles ausgeben

	fmt.Println("Dies ist der Text der Webseite")
	fmt.Println(htmlDoc.Text())

	fmt.Println("Bitte einen Suchbegriff eingeben:")
	var searchquery string
	fmt.Scanln(&searchquery)

	docContainsSearchQuery := htmlDoc.Contains(searchquery)
	if docContainsSearchQuery {
		fmt.Println("Der Suchbegriff ist enthalten.")
	} else {
		fmt.Println("Der Suchbegriff ist nicht enthalten.")
	}
}
