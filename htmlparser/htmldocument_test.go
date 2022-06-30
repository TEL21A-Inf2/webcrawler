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
	// [link text: url]
}

func ExampleHtmlDocument_Links_pageWithOneLinkFromFile() {
	doc, _ := FromFile("testdata/simplepageonelink.html")
	fmt.Println(doc.Links())

	// Output:
	// [link text: url]
}

func ExampleHtmlDocument_Links_pageWithTwoLinksFromString() {
	doc1, _ := FromString(testdata.SimplePageTwoLinks)
	fmt.Println(doc1.Links())

	// Output:
	// [link text 1: url1 link text 2: url2]
}

func ExampleHtmlDocument_Links_pageWithTwoLinksFromFile() {
	doc, _ := FromFile("testdata/simplepagetwolinks.html")
	fmt.Println(doc.Links())

	// Output:
	// [link text 1: url1 link text 2: url2]
}

func ExampleHtmlDocument_Links_pageWithTwoLinksAndTextFromString() {
	doc1, _ := FromString(testdata.SimplePageTwoLinksAndText)
	fmt.Println(doc1.Links())
	fmt.Println(doc1.Text())

	// Output:
	// [link text 1: url1 link text 2: url2]
	// My First Heading
	// My first paragraph.
	// Lorem ipsum
	// link text 1
	// link text 2
	// dolor sit amet
}

func ExampleHtmlDocument_Links_pageWithTwoLinksAndTextFromFile() {
	doc, _ := FromFile("testdata/simplepagetwolinksandtext.html")
	fmt.Println(doc.Links())
	fmt.Println(doc.Text())

	// Output:
	// [link text 1: url1 link text 2: url2]
	// My First Heading
	// My first paragraph.
	// Lorem ipsum
	// link text 1
	// link text 2
	// dolor sit amet
}
