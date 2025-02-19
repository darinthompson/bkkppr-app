package main

import (
	"log"

	"github.com/darinthompson/bkkppr-app/internal/config"
	"github.com/darinthompson/bkkppr-app/internal/handlers"
	"github.com/darinthompson/bkkppr-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	repository.ConnectDatabase(cfg)

	// Create Gin router
	r := gin.Default()

	// Define routes
	r.POST("/register", handlers.CreateUserHandler)
	r.GET("/users", handlers.GetUsersHandler)
	r.GET("/users/:id", handlers.GetUserByID)

	r.POST("/books", handlers.CreateBookhandler)
	r.GET("/books", handlers.GetBooksHandler)
	r.GET("/user/books", handlers.GetAll)

	// Start server
	log.Println("Server is running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
