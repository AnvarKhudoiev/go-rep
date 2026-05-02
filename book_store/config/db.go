package config

import (
	"ginExample/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=Anvar2005! dbname=book_store port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate your models
	database.AutoMigrate(&models.Author{}, &models.Category{}, &models.Book{}, &models.User{}, &models.FavoriteBook{})

	DB = database
}
