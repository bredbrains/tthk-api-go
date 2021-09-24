package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          int `gorm:"PrimaryKey"`
	VKontakteID int
	TelegramID  int
}
