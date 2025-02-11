package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Author string `json:"author,omitempty"`
	ISBN   string `gorm:"unique" json:"isbn,omitempty"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"owner"`

	Genre        string `gorm:"size:255;not null"`
	Publish_Date string `gorm:"size:255;not null"`
}
