package repository

import (
	"log"

	"github.com/darinthompson/bkkppr-app/internal/config"
	"github.com/darinthompson/bkkppr-app/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase(cfg *config.Config) {
	dsn := cfg.DatabaseURL // Fetch DB connection string from config
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// AutoMigrate to create tables
	db.AutoMigrate(&models.User{}, &models.Book{})

	DB = db
}
