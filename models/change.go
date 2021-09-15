package models

import (
	"time"
)

type Change struct {
	Date    time.Time
	Group   string
	Lessons string
	Teacher string
	Room    string
}
