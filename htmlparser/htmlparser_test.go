package htmlparser

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

func ExampleGetLinks_PageWithoutLinks() {
	fmt.Println(GetLinks(simplePageNoLinks))

	// Output:
	// []
}

func ExampleGetLinks_PageWithOneLink() {
	fmt.Println(GetLinks(simplePageOneLink))

	// Output:
	// [url]
}

func ExampleGetLinks_PageWithTwoLinks() {
	fmt.Println(GetLinks(simplePageTwoLinks))

	// Output:
	// [url1 url2]
}
