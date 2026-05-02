package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories retrieves all categories from the database.
func GetCategories(c *gin.Context) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// AddCategory adds a new category to the database.
func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newCategory.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	result := config.DB.Create(&newCategory)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
