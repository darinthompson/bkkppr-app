package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetUserByID(c *gin.Context) {
	userID := c.Param("id") // Extract user ID from URL

	// Convert string userID to uint
	parsedID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call repository function
	user, err := repository.GetUserByID(uint(parsedID))
	fmt.Println(parsedID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user as JSON
	c.JSON(http.StatusOK, user)
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

func GetAll(c *gin.Context) {
	users, err := repository.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch users and books"})
		return
	}
	c.JSON(http.StatusOK, users)
}
