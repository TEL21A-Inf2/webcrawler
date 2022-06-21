package htmldocument

import "fmt"

const simplePageNoLinks = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>
`

const simplePageOneLink = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

<a href="url">link text</a>

</body>
</html>
`

const simplePageTwoLinks = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

<a href="url1">link text 1</a>
<a href="url2">link text 2</a>

</body>
</html>
`

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
