package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Query(htmlContent string, selector string) *goquery.Document {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	LogError(err)
	return dom
}
