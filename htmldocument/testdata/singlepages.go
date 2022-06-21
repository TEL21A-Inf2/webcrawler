package testdata

const (
	SimplePageNoLinks = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>
`

	SimplePageOneLink = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

<a href="url">link text</a>

</body>
</html>
`

	SimplePageTwoLinks = `
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
)
