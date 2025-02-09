package repository

import (
	"github.com/darinthompson/bkkppr-app/internal/models"
)

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	return users, err
}
