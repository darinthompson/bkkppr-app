package repository

import (
	"errors"
	"log"

	"github.com/darinthompson/bkkppr-app/internal/models"
	"github.com/darinthompson/bkkppr-app/internal/utils"
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

func Login(authInput models.AuthInput) (*models.User, error) {
	var user models.User
	err := DB.Preload("Books").First(&user, "username = ? ", authInput.Username).Error
	if err != nil {
		log.Printf("ERROR FETCHING USER WITH USERNAME: %v: %v", authInput.Username, err)
		return nil, err
	}

	if !utils.CheckPasswordHash(authInput.Password, user.Password) {
		log.Println("INVALID PASSWORD")
		return nil, errors.New("invalid username or password")
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetAll() ([]models.User, error) {
	var users []models.User
	err := DB.Model(&models.User{}).Preload("Books").Find(&users).Error
	return users, err
}
