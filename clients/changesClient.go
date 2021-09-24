package clients

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bredbrains/tthk-api-go/models"
	"github.com/bredbrains/tthk-api-go/models/enums"
	"github.com/bredbrains/tthk-api-go/models/enums/estonian"
	"github.com/bredbrains/tthk-api-go/models/enums/selectors"
	"github.com/bredbrains/tthk-api-go/modules"
	"github.com/bredbrains/tthk-api-go/utils/globals"
	"strings"
	"time"
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
	document.Find(selectors.Table).Each(c.processChangeTable)
}

func (c *ChangesClient) processChangeTable(_ int, table *goquery.Selection) {
	cells := table.Find(selectors.TableRow)
	cells.Each(c.processChangeTableRow)
}

func (c *ChangesClient) processChangeTableRow(_ int, tableRow *goquery.Selection) {
	change := models.Change{}
	if tableRow.Find(selectors.TableCell).Eq(1).Text() == "Kuupäev" {
		return
	}
	tableRow.Find(selectors.TableCell).Each(func(index int, tableCell *goquery.Selection) {
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
	if text == estonian.DroppedOut {
		change.Status = enums.DroppedOut
	} else if text == estonian.Lunch {
		change.Status = enums.Lunch
	} else if text == estonian.Homework {
		change.Status = enums.Homework
	} else if text == estonian.Scheduled {
		change.Status = enums.Scheduled
	} else {
		change.Teacher = text
	}
	return change
}
