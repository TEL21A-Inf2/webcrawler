package htmlparser

import (
	"fmt"
	"net/url"
)

// Repräsentiert einen Hyperlink bestehend aus der Url und dem Link-Text.
type Hyperlink struct {
	Url  *url.URL
	Text string
}

func NewHyperlink(rawUrl, text string) Hyperlink {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return Hyperlink{nil, ""}
	}
	return Hyperlink{parsedUrl, text}
}

func (link Hyperlink) String() string {
	return fmt.Sprintf("%s: %s", link.Text, link.Url.String())
}

func (link Hyperlink) IsValid() bool {
	return link.Url != nil
}
