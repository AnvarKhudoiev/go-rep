package main

import (
	"fmt"
	"log"

	"Todo_Manager/todo_service/config"
	"Todo_Manager/todo_service/clients"
	"Todo_Manager/todo_service/models"
	"Todo_Manager/todo_service/handlers"
	"Todo_Manager/todo_service/middleware"

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

	r.POST("/internal/categories/defaults", func(c *gin.Context) {
		var req struct {
			UserID uint `json:"user_id"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		defaultCategories := []models.Category{
			{Name: "Личное", Color: "#4CAF50", UserID: req.UserID},
			{Name: "Работа", Color: "#F44336", UserID: req.UserID},
		}

		if err := config.DB.Create(&defaultCategories).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create categories"})
			return
		}

		c.JSON(201, gin.H{"message": "Categories created"})
	})

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) 
	{
		
		api.POST("/todos", handlers.CreateTodo)
		api.GET("/todos", handlers.GetTodos)
		api.GET("/todos/:id", handlers.GetTodoByID)
		api.PUT("/todos/:id", handlers.UpdateTodo)
		api.DELETE("/todos/:id", handlers.DeleteTodo)

		api.POST("/categories", handlers.CreateCategory)
		api.GET("/categories", handlers.GetCategories)
	}


	// ВРЕМЕННЫЙ РОУТ ДЛЯ ТЕСТА RESTY
	r.GET("/test-resty/:user_id", func(c *gin.Context) {
		userIDStr := c.Param("user_id")
		
		// Конвертируем строку в число (для простоты теста используем хак, 
		// в реальном коде лучше через strconv.Atoi)
		var userID uint
		fmt.Sscanf(userIDStr, "%d", &userID)

		// Вызываем наш клиент!
		user, err := clients.GetUserByID(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Ура! Связь работает!",
			"user": user,
		})
	})

	// 5. Запускаем сервис на порту 8082
	log.Println("TodoService is running on port 8082...")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Failed to start TodoService:", err)
	}
}