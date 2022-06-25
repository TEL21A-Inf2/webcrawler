package httpreader

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GoQueryDocumentFromBytes(data []byte) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(bytes.NewReader(data))
}

func GoQueryDocumentFromString(data string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(data))
}

func GoQueryDocumentFromFile(path string) (*goquery.Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return goquery.NewDocumentFromReader(bufio.NewReader(file))
}

func GoQueryDocumentFromUrl(url string) (*goquery.Document, error) {
	data, err := Get(url)
	if err != nil {
		return nil, err
	}
	return GoQueryDocumentFromBytes(data)
}
