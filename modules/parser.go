package modules

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

const tableSelector string = "table"

// Wrapper for scraping web pages with internal logic
func Scrape(url string) *goquery.Document {
	var err error
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return document
}
