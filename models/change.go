package models

import (
	"time"
)

type Change struct {
	Date    time.Time `json:"date"`
	Group   string    `json:"group"`
	Lessons string    `json:"lessons"`
	Teacher string    `json:"teacher"`
	Room    string    `json:"room,omitempty"`
}

func (c *Change) Validate() bool {
	if c.Group == "" {
		return false
	} else if c.Lessons == "" {
		return false
	} else if c.Teacher == "" {
		return false
	} else if time.Now().Year() != c.Date.Year() {
		return false
	}
	return true
}
