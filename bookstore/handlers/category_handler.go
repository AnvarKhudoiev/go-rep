package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"bookstore/models"
)

var Categories []models.Category
var categoryID = 1

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, Categories)
}

func CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.BindJSON(&cat); err != nil || cat.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	cat.ID = categoryID
	categoryID++
	Categories = append(Categories, cat)
	c.JSON(http.StatusCreated, cat)
}