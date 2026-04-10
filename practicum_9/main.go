package main

import (
	"github.com/gin-gonic/gin"
	"practicum_9/config"
	"practicum_9/handlers"
	"practicum_9/middleware"
)

func main() {
	r := gin.Default()
	config.ConnectDB()


	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/todos", handlers.CreateTodo)
		protected.GET("/todos", handlers.GetTodos)
		protected.GET("/todos/:id", handlers.GetTodoById)
		protected.PUT("/todos/:id", handlers.UpdateTodo)
		protected.DELETE("/todos/:id", handlers.DeleteTodo)
		protected.POST("/authors", handlers.CreateAuthor)
		protected.GET("/authors", handlers.GetAuthors)
		protected.GET("/authors/:id", handlers.GetAuthorById)
		protected.POST("/authors/:id/todos", handlers.CreateTodoForAuthor)
	}

	r.Run(":8080")

}
