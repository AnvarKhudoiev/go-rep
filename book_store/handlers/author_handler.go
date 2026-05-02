package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuthors retrieves all authors from the database.
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	result := config.DB.Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// AddAuthor adds a new author to the database.
func AddAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newAuthor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author name is required"})
		return
	}

	result := config.DB.Create(&newAuthor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add author"})
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
}
