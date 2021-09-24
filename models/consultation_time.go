package models

import "time"

type ConsultationTime struct {
	weekday        time.Weekday
	startTime      time.Time
	endTime        time.Time
	additionalData string
}
