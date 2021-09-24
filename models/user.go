package models

import "gorm.io/gorm"

// User Model to submit
type User struct {
	gorm.Model
	ID          int `gorm:"PrimaryKey"`
	VKontakteID int
	TelegramID  int
	DiscordID   int
	Group       string
}
