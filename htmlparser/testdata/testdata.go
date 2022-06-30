package testdata

import _ "embed"

//go:embed simplepagenolinks.html
var SimplePageNoLinks string

//go:embed simplepageonelink.html
var SimplePageOneLink string

//go:embed simplepagetwolinks.html
var SimplePageTwoLinks string

//go:embed simplepagetwolinksandtext.html
var SimplePageTwoLinksAndText string
