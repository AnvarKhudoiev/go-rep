package main

import (
	"practicum_8/config"
	"practicum_8/handlers"
	"practicum_8/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Author{}, &models.Todo{})
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodoById)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	r.POST("/authors", handlers.CreateAuthor)
	r.GET("/authors", handlers.GetAuthors)
	r.GET("/authors/:id", handlers.GetAuthorById)
	r.POST("/authors/:id/todos", handlers.CreateTodoForAuthor)
	r.Run(":8080")
}