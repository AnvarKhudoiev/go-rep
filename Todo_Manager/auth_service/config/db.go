package config

import (
	"fmt"
	"log"
	"os"

	// Замени "Todo_Manager/auth_service/models" на реальный путь к твоим моделям
	"Todo_Manager/auth_service/models" 

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error:", err)
		os.Exit(1)
	}

	// АВТОМИГРАЦИЯ: GORM прочитает твои структуры в Go и создаст таблицы
	// Добавь сюда все свои модели через запятую
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	DB = db
	log.Println("Database connection established and tables migrated!")
}