package repository

import (
	"log"

	"github.com/darinthompson/bkkppr-app/internal/models"
)

func CreateBook(book *models.Book) error {
	return DB.Create(book).Error
}

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := DB.Find(&books).Error
	return books, err
}

func GetBooksByUserId(userID uint) ([]models.Book, error) {
	var books []models.Book

	// Execute the query
	err := DB.Where("user_id = ?", userID).Find(&books).Error
	if err != nil {
		log.Printf("ERROR FETCHING BOOKS WITH USERID: %d - %v", userID, err)
		return nil, err
	}

	return books, nil
}
