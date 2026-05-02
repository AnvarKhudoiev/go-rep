package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books from the database with pagination
func GetBooks(c *gin.Context) {
	// Get page and limit from query parameters, set defaults if not provided
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Convert string to integer
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Calculate offset
	offset := (pageInt - 1) * limitInt

	var books []models.Book
	var total int64

	// Get total count of books
	config.DB.Model(&models.Book{}).Count(&total)

	// Get books with pagination, without loading relationships
	result := config.DB.Model(&models.Book{}).
		Select("id, title, author_id, category_id").
		Limit(limitInt).
		Offset(offset).
		Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       books,
		"total":      total,
		"page":       pageInt,
		"limit":      limitInt,
		"totalPages": int(math.Ceil(float64(total) / float64(limitInt))),
	})
}

// AddBook adds a new book to the database
func AddBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newBook.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book title is required"})
		return
	}

	result := config.DB.Create(&newBook)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	// Fetch the complete book data with related entities
	config.DB.Preload("Author").Preload("Category").First(&newBook, newBook.ID)
	c.JSON(http.StatusCreated, newBook)
}

// GetBookByID retrieves a single book by its ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	result := config.DB.Preload("Author").Preload("Category").First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book by its ID
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	// Check if book exists
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Bind new data
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the book
	result := config.DB.Save(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Return updated book with relationships
	config.DB.Preload("Author").Preload("Category").First(&book, id)
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Book{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
