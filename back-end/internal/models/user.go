package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	Username  string `json:"username" gorm:"not null;unique"`
	Password  string `json:"password"`
	Books     []Book `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"books"`
}
