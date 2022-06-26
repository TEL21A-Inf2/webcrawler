package htmlparser

import (
	"fmt"

	"github.com/tel21a-inf2/webcrawler/htmlparser/testdata"
)

func ExampleHtmlDocument_Links_pageWithoutLinksFromString() {
	doc, _ := FromString(testdata.SimplePageNoLinks)
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleHtmlDocument_Links_pageWithoutLinksFromFile() {
	doc, _ := FromFile("testdata/simplepagenolinks.html")
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleHtmlDocument_Links_pageWithOneLinkFromString() {
	doc, _ := FromString(testdata.SimplePageOneLink)
	fmt.Println(doc.Links())

	// Output:
	// [{url link text}]
}

func ExampleHtmlDocument_Links_pageWithOneLinkFromFile() {
	doc, _ := FromFile("testdata/simplepageonelink.html")
	fmt.Println(doc.Links())

	// Output:
	// [{url link text}]
}

func ExampleHtmlDocument_Links_pageWithTwoLinksFromString() {
	doc1, _ := FromString(testdata.SimplePageTwoLinks)
	fmt.Println(doc1.Links())

	// Output:
	// [{url1 link text 1} {url2 link text 2}]
}

func ExampleHtmlDocument_Links_pageWithTwoLinksFromFile() {
	doc, _ := FromFile("testdata/simplepagetwolinks.html")
	fmt.Println(doc.Links())

	// Output:
	// [{url1 link text 1} {url2 link text 2}]
}
