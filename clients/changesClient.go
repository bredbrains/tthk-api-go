package clients

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
	"tthkAPI/models"
	"tthkAPI/modules"
	"tthkAPI/utils/globals"
)

type ChangesClient struct{}

const ChangesUrl = globals.Changes

func (c *ChangesClient) Parse() {
	document := modules.Scrape(ChangesUrl)
	processChangesTables(document)
}

func processChangesTables(document *goquery.Document) {
	document.Find("table").Each(processChangeTable)
}

func processChangeTable(index int, table *goquery.Selection) {
	table.Find("tr").Each(processChangeTableRow)
}

func processChangeTableRow(index int, tableRow *goquery.Selection) {
	change := models.Change{}
	tableRow.Find("td").Each(func(index int, tableCell *goquery.Selection) {
		var err error
		cellText := tableCell.Text()
		switch index {
		case 1:
			change.Date, err = time.Parse(time.Layout, cellText)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}
