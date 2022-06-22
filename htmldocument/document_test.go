package htmldocument

import (
	"fmt"

	"github.com/tel21a-inf2/webcrawler/htmldocument/testdata"
)

func ExampleDocument_Links_PageWithoutLinks_FromString() {
	doc, _ := FromString(testdata.SimplePageNoLinks)
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleDocument_Links_PageWithoutLinks_FromFile() {
	doc, _ := FromFile("testdata/simplepagenolinks.html")
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleDocument_Links_PageWithOneLink_FromString() {
	doc, _ := FromString(testdata.SimplePageOneLink)
	fmt.Println(doc.Links())

	// Output:
	// [url]
}

func ExampleDocument_Links_PageWithOneLink_FromFile() {
	doc, _ := FromFile("testdata/simplepageonelink.html")
	fmt.Println(doc.Links())

	// Output:
	// [url]
}

func Exampledocument_Links_PageWithTwoLinks_FromString() {
	doc, _ := FromString(testdata.SimplePageTwoLinks)
	fmt.Println(doc.Links())

	// Output:
	// [url1 url2]
}

func ExampleDocument_Links_PageWithTwoLinks_FromFile() {
	doc, _ := FromFile("testdata/simplepagetwolinks.html")
	fmt.Println(doc.Links())

	// Output:
	// [url1 url2]
}
