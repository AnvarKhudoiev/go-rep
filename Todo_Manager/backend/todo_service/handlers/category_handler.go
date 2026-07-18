package handlers

import (
	"backend/todo_service/config"
	"backend/todo_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = userID
	input.IsDefault = false

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func GetCategories(c *gin.Context) {
	userID := c.GetUint("user_id")
	var categories []models.Category

	// Возвращаем дефолтные (user_id=0) + категории пользователя
	if err := config.DB.
		Where("is_default = ? OR user_id = ?", true, userID).
		Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")
	var category models.Category

	if err := config.DB.
		Where("id = ? AND (is_default = ? OR user_id = ?)", id, true, userID).
		First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")
	var category models.Category

	if err := config.DB.Where("id = ? AND user_id = ? AND is_default = ?", id, userID, false).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found or cannot edit default"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	result := config.DB.Where("id = ? AND user_id = ? AND is_default = ?", id, userID, false).Delete(&models.Category{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found or cannot delete default"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}