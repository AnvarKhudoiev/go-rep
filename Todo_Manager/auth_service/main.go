package main

import (
	"log"

	"Todo_Manager/auth_service/config"
	"Todo_Manager/auth_service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Error loading .env file. Using system environment variables.")
	}
	config.ConnectDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users/:id", handlers.GetUserByID)

	log.Println("AuthService is running on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start AuthService:", err)
	}
}
