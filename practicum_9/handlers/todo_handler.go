package handlers

import (
	"net/http"
	"practicum_9/config"
	"practicum_9/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Preload("Author").Find(&todos) // загружаем автора
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.Preload("Author").First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Todo{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
