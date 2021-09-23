package clients

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
	"tthkAPI/models"
	"tthkAPI/models/enums"
	"tthkAPI/models/enums/estonianEnums"
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

func (c *ChangesClient) processChangeTable(_ int, table *goquery.Selection) {
	cells := table.Find("tr")
	cells.Each(c.processChangeTableRow)
}

func (c *ChangesClient) processChangeTableRow(_ int, tableRow *goquery.Selection) {
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
			change = determineStatusText(change, cellText)
		case 5:
			change.Room = cellText
		}
		return
	})
	if change.Validate() {
		c.changes = append(c.changes, change)
	}
}

func determineStatusText(change models.Change, text string) models.Change {
	text = strings.TrimSpace(text)
	if text == estonianEnums.DroppedOut {
		change.Status = enums.DroppedOut
	} else if text == estonianEnums.Lunch {
		change.Status = enums.Lunch
	} else if text == estonianEnums.Homework {
		change.Status = enums.Homework
	} else if text == estonianEnums.Scheduled {
		change.Status = enums.Scheduled
	} else {
		change.Teacher = text
	}
	return change
}
