package models

import (
	"time"
)

type Change struct {
	date    time.Time
	group   string
	lessons string
	teacher string
	room    string
}
