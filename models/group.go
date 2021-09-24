package models

import "github.com/bredbrains/tthk-api-go/models/enums"

type Group struct {
	code       string
	department enums.Department
	language   enums.Language
	teacher    Teacher
}
