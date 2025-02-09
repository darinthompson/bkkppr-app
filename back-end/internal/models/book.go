package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title        string `gorm:"size:255;not null"`
	UserId       string
	Author       string `gorm:"size:255;not null"`
	Genre        string `gorm:"size:255;not null"`
	Publish_Date string `gorm:"size:255;not null"`
}
