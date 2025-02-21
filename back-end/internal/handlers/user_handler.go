package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/darinthompson/bkkppr-app/internal/models"
	"github.com/darinthompson/bkkppr-app/internal/repository"
	"github.com/darinthompson/bkkppr-app/internal/utils"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles creating a new user
func CreateUserHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func Login(c *gin.Context) {
	var authBody = models.AuthInput{}

	if c.Bind(&authBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	user, err := repository.Login(authBody)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Println("ERROR: SECRET_KEY environment variable is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating Token"})
		return
	}

	fmt.Println(tokenString, err)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
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

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
