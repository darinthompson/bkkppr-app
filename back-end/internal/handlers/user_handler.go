package handlers

import (
	"net/http"

	"github.com/darinthompson/bkkppr-app/internal/models"
	"github.com/darinthompson/bkkppr-app/internal/repository"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles creating a new user
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// GetUsersHandler handles retrieving all users
func GetUsersHandler(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
