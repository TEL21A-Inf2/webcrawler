package htmldocument

import "fmt"

func ExampleDocument_Links_PageWithoutLinks() {
	doc, _ := FromString(simplePageNoLinks)
	fmt.Println(doc.Links())

	// Output:
	// []
}

func ExampleDocument_Links_PageWithOneLink() {
	doc, _ := FromString(simplePageOneLink)
	fmt.Println(doc.Links())

	// Output:
	// [url]
}

func Exampledocument_Links_PageWithTwoLinks() {
	doc, _ := FromString(simplePageTwoLinks)
	fmt.Println(doc.Links())

	// Output:
	// [url1 url2]
}
