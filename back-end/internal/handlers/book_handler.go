package handlers

import (
	"net/http"
	"strconv"

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

	// userInterface, exists := c.Get("user")
	// if !exists {
	// 	c.JSON()
	// }

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

func GetBooksByUserId(c *gin.Context) {
	userIDParam := c.Param("id")

	// Convert ID from string to uint
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch books using repository function
	books, err := repository.GetBooksByUserId(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	// Return books as JSON
	c.JSON(http.StatusOK, books)
}
