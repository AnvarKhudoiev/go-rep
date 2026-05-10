package main

import (
	"log"

	"backend/auth_service/config"
	"backend/auth_service/handlers"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users/:id", handlers.GetUserByID)

	log.Println("AuthService is running on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start AuthService:", err)
	}
}
