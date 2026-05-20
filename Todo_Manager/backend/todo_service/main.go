package main

import (
	"fmt"
	"log"

	"backend/todo_service/clients"
	"backend/todo_service/config"
	"backend/todo_service/handlers"
	"backend/todo_service/middleware"
	"backend/todo_service/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func seedDefaultCategories() {
	defaults := []models.Category{
		{Name: "Личное",   Color: "#4CAF50", IsDefault: true},
		{Name: "Работа",   Color: "#F44336", IsDefault: true},
		{Name: "Учёба",    Color: "#2196F3", IsDefault: true},
		{Name: "Здоровье", Color: "#FF9800", IsDefault: true},
		{Name: "Покупки",  Color: "#9C27B0", IsDefault: true},
	}

	for _, cat := range defaults {
		var existing models.Category
		result := config.DB.Where("name = ? AND is_default = ?", cat.Name, true).First(&existing)
		if result.RowsAffected == 0 {
			config.DB.Create(&cat)
		}
	}
}

func main() {
	err := godotenv.Load("todo_service/.env")
	if err != nil {
		log.Println("Warning: Error loading .env file. Using system environment variables.")
	}

	config.ConnectDB()
	seedDefaultCategories()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Todos — overdue ДОЛЖЕН быть до /:id иначе gin матчит "overdue" как id
		api.GET("/todos/overdue", handlers.GetOverdueTodos)
		api.POST("/todos", handlers.CreateTodo)
		api.GET("/todos", handlers.GetTodos)
		api.GET("/todos/:id", handlers.GetTodoByID)
		api.PUT("/todos/:id", handlers.UpdateTodo)
		api.DELETE("/todos/:id", handlers.DeleteTodo)

		// Categories — полный CRUD
		api.POST("/categories", handlers.CreateCategory)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategoryByID)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)

		// Stats
		api.GET("/stats", handlers.GetStats)
	}


	r.GET("/test-resty/:user_id", func(c *gin.Context) {
		var userID uint
		fmt.Sscanf(c.Param("user_id"), "%d", &userID)
		user, err := clients.GetUserByID(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Связь работает!", "user": user})
	})

	log.Println("TodoService is running on port 8082...")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Failed to start TodoService:", err)
	}
}