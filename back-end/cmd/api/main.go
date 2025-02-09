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
	r.POST("/users", handlers.CreateUserHandler)
	r.GET("/users", handlers.GetUsersHandler)

	// Start server
	log.Println("Server is running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
