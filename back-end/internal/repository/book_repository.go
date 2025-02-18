package repository

import (
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
