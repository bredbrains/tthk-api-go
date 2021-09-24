package models

import "github.com/bredbrains/tthk-api-go/models/enums"

type Teacher struct {
	name       string
	room       string
	email      string
	department enums.Department
	times      []ConsultationTime
}
