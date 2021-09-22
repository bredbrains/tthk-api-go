package clients

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
	"tthkAPI/models"
	"tthkAPI/modules"
	"tthkAPI/utils/globals"
)

type ChangesClient struct {
	changes []models.Change
}

const ChangesUrl = globals.Changes

func (c *ChangesClient) Parse() []models.Change {
	c.changes = []models.Change{}
	document := modules.Scrape(ChangesUrl)
	c.processChangesTables(document)
	return c.changes
}

func (c *ChangesClient) processChangesTables(document *goquery.Document) {
	document.Find("table").Each(c.processChangeTable)
}

func (c *ChangesClient) processChangeTable(index int, table *goquery.Selection) {
	cells := table.Find("tr")
	cells.Each(c.processChangeTableRow)
}

func (c *ChangesClient) processChangeTableRow(index int, tableRow *goquery.Selection) {
	change := models.Change{}
	if tableRow.Find("td").Eq(1).Text() == "Kuupäev" {
		return
	}
	tableRow.Find("td").Each(func(index int, tableCell *goquery.Selection) {
		var err error
		cellText := strings.TrimSpace(tableCell.Text())
		switch index {
		case 1:
			change.Date, err = time.Parse("02.01.2006", cellText)
			if err != nil {
				return
			}
		case 2:
			change.Group = cellText
		case 3:
			change.Lessons = cellText
		case 4:
			change.Teacher = cellText
		case 5:
			change.Room = cellText
		}
		return
	})
	if change.Validate() {
		c.changes = append(c.changes, change)
	}
}
