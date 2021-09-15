package models

import "tthkAPI/models/enums"

type Teacher struct {
	name       string
	room       string
	email      string
	department enums.Department
	times      []ConsultationTime
}
