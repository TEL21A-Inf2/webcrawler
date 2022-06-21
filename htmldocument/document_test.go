package htmldocument

import (
	"fmt"

	"github.com/tel21a-inf2/webcrawler/htmldocument/testdata"
)

func ExampleDocument_Links_PageWithoutLinks() {
	doc, _ := FromString(testdata.SimplePageNoLinks)
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleDocument_Links_PageWithOneLink() {
	doc, _ := FromString(testdata.SimplePageOneLink)
	fmt.Println(doc.Links())

	// Output:
	// [url]
}

func Exampledocument_Links_PageWithTwoLinks() {
	doc, _ := FromString(testdata.SimplePageTwoLinks)
	fmt.Println(doc.Links())

	// Output:
	// [url1 url2]
}
