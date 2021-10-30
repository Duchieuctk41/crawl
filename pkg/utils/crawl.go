package utils

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func CrawlPage(link string) int {
	selector := ".titlePagi"
	sel := ".titlePagi"

	// lay tong so trang va cac san pham
	pageNumber, err := CrawlText(link, selector, sel)
	LogError(err)

	if pageNumber != "" {
		splitted := strings.Split(pageNumber, " ")
		page, _ := strconv.Atoi(splitted[1])
		return page
	}
	return 0
}

func CrawlText(url string, selector string, sel interface{}) (string, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	var page string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Text(sel, &page),
	)
	LogError(err)

	return page, nil
}

// crawl product
func CrawlHTML(url string, selector string, sel interface{}) (string, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &html),
	)
	LogError(err)

	return html, nil
}
