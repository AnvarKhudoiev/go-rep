package main

import (
	"ginExample/config"
	"ginExample/handlers"
	"ginExample/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/books", handlers.GetBooks)
		protected.POST("/books", handlers.AddBook)
		protected.GET("/books/:id", handlers.GetBookByID)
		protected.PUT("/books/:id", handlers.UpdateBook)
		protected.DELETE("/books/:id", handlers.DeleteBook)

		protected.GET("/authors", handlers.GetAuthors)
		protected.POST("/authors", handlers.AddAuthor)

		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)

		protected.GET("/favorites", handlers.GetFavoriteBooks)
		protected.PUT("/:bookId/favorites", handlers.AddFavorite)
		protected.DELETE("/:bookId/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
