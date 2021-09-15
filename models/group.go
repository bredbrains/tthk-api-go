package models

import "tthkAPI/models/enums"

type Group struct {
	code       string
	department enums.Department
	language   enums.Language
	teacher    Teacher
}
