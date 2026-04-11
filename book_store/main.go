package main

import (
	"ginExample/config"
	"ginExample/handlers"
	"ginExample/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	config.ConnectDatabase()

	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Book routes
		protected.GET("/books", handlers.GetBooks)
		protected.POST("/books", handlers.AddBook)
		protected.GET("/books/:id", handlers.GetBookByID)
		protected.PUT("/books/:id", handlers.UpdateBook)
		protected.DELETE("/books/:id", handlers.DeleteBook)

		// Author routes
		protected.GET("/authors", handlers.GetAuthors)
		protected.POST("/authors", handlers.AddAuthor)

		// Category routes
		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)

		// Fav routes
		protected.GET("/favorites", handlers.GetFavoriteBooks)
		protected.PUT("/:bookId/favorites", handlers.AddFavorite)
		protected.DELETE("/:bookId/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
