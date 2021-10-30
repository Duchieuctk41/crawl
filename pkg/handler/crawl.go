package handler

import (
	"crawl/pkg/model"
	"crawl/pkg/service"
	"crawl/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Handler struct {
	Service service.IService
}

type IHandler interface {
	CrawlCollection(w http.ResponseWriter, r *http.Request)
}

func NewHandler(srv service.IService) IHandler {
	return &Handler{
		Service: srv,
	}
}

// lay cac collection
func (h *Handler) CrawlCollection(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnHTML(".menu-sub_hasSub_lv1", func(e *colly.HTMLElement) {
		name := e.Attr("data-text")
		collect := model.Collection{}
		if name != "" {
			collect.Name = name
		}

		e.ForEach(".menu-sub_hasSub_lv2 > li > a", func(_ int, el *colly.HTMLElement) {
			col2 := model.Level1{
				Name: el.Text,
				Link: el.Attr("href"),
			}
			if col2.Link == "/" {
				return
			}
			collect.Level1 = append(collect.Level1, col2)
		})

		url := "https://www.maisononline.vn"
		if collect.Name != "" {
			link := url + collect.Level1[0].Link
			page := utils.CrawlPage(link)
			if page != 0 {
				collect.Page = page
				err := h.Service.CreateCollection(&collect)
				if err != nil {
					fmt.Fprintf(w, "success")
					return
				}

				CrawlProductLink(link, page, h)
			}

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.maisononline.vn/")
}

func CrawlProductLink(link string, page int, h *Handler) {
	selector := ".titlePagi"
	sel := ".product-list-collection"
	for i := 1; i <= page; i++ {
		url := link + "?page=" + strconv.Itoa(i)
		products, err := utils.CrawlHTML(url, selector, sel)
		utils.LogError(err)
		GetLinkProduct(products, ".product-list-collection", h)
	}
}

func GetLinkProduct(htmlContent string, selector string, h *Handler) {
	dom := utils.Query(htmlContent, selector)
	sTitle := ".product-loop-inner"
	dom.Find(sTitle).Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Find(".product-loop-img > a").Attr("href")
		url := "https://www.maisononline.vn" + link
		CrawlProductDetail(url, h)
	})
}

// crawl detail product
func CrawlProductDetail(link string, h *Handler) {
	selector := ".image img"
	sel := ".product-content"
	htmlDetail, err := utils.CrawlHTML(link, selector, sel)
	utils.LogError(err)
	GetDetailProduct(htmlDetail, ".product-content", h)
}

// crawl detail product
func GetDetailProduct(htmlContent string, selector string, h *Handler) {
	dom := utils.Query(htmlContent, selector)
	details := []string{}
	sizes := []string{}
	colors := []string{}
	imgs := []string{}
	totalColor := 0

	sku := strings.TrimSpace(dom.Find(".details__sku").Text())
	title := strings.TrimSpace(dom.Find(".details__prd-title").Text())
	vendor := strings.TrimSpace(dom.Find(".details__brand").Text())
	price := dom.Find(".price").Text()
	initialPrice := dom.Find(".compare").Text()

	sale := strings.TrimSpace(dom.Find(".sale").Text())
	dom.Find("#product-content ul li").Each(func(i int, selection *goquery.Selection) {
		details = append(details, selection.Text())
	})
	dom.Find(".details__sizes_data .size").Each(func(i int, selection *goquery.Selection) {
		size, _ := selection.Attr("data-size")
		sizes = append(sizes, size)
	})
	dom.Find(".image").Each(func(i int, selection *goquery.Selection) {
		color, _ := selection.Attr("data-color")
		img, _ := selection.Find("img").Attr("src")
		if color != "" {
			colors = append(colors, color)
			totalColor++
		}
		imgs = append(imgs, img)
	})

	item := model.Item{
		Title:        title,
		Vendor:       vendor,
		Sale:         sale,
		Price:        price,
		InitialPrice: initialPrice,
		TotalColor:   strconv.Itoa(totalColor),
		Sku:          sku,
		Images:       imgs,
		Colors:       colors,
		Sizes:        sizes,
		Details:      details,
	}

	err := h.Service.CreateProduct(&item)
	if err != nil {
		return
	}
}
