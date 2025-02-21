package main

import (
	"log"

	"github.com/darinthompson/bkkppr-app/internal/config"
	"github.com/darinthompson/bkkppr-app/internal/handlers"
	"github.com/darinthompson/bkkppr-app/internal/middleware"
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
	r.POST("/signup", handlers.CreateUserHandler)
	r.POST("/login", handlers.Login)
	r.GET("/validate", middleware.RequireAuth, handlers.Validate)
	r.GET("/users", handlers.GetUsersHandler)
	r.GET("/users/:id", handlers.GetUserByID)

	r.POST("/books", handlers.CreateBookhandler)
	r.GET("/books", handlers.GetBooksHandler)

	// Start server
	log.Println("Server is running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
