package handlers

import (
	"net/http"

	"github.com/darinthompson/bkkppr-app/internal/models"
	"github.com/darinthompson/bkkppr-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func CreateBookhandler(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}

	book.UserID = 1

	if err := repository.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Book created Successfully", "Book": book})
}

func GetBooksHandler(c *gin.Context) {
	books, err := repository.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// func GetBooksByUserID(c *gin.Context) {
// 	books, err models.Book
// }
