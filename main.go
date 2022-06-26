package main

import (
	"fmt"
	"log"

	"github.com/tel21a-inf2/webcrawler/htmlparser"
)

// Ein einfaches Programm, das den Benutzer nach einer URL fragt und die Links auf
// auf dieser Seite ausgibt.

func main() {
	fmt.Println("Bitte eine URL eingeben:")
	var url string
	fmt.Scanln(&url)

	htmlDoc, err := htmlparser.FromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Es wurden folgende Links gefunden:")
	fmt.Println(htmlDoc.Links())
}
