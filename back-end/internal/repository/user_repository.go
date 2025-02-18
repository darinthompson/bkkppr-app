package repository

import (
	"log"

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

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := DB.Preload("Books").First(&user, userID).Error
	if err != nil {
		log.Printf("ERROR FETCHING USER WITH ID: %d: %v", userID, err)
		return nil, err
	}
	return &user, nil
}

func GetAll() ([]models.User, error) {
	var users []models.User
	err := DB.Model(&models.User{}).Preload("Books").Find(&users).Error
	return users, err
}
