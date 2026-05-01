package handlers

import (
	"Todo_Manager/todo_service/config"
	"Todo_Manager/todo_service/models"
	"Todo_Manager/todo_service/clients"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	userID := c.GetUint("user_id")

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = userID

	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	config.DB.Preload("Category").Preload("Tag").First(&todo, todo.ID)
	c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	userID := c.GetUint("user_id")
	var todos []models.Todo
	if err := config.DB.
		Preload("Category").
		Preload("Tags").
		Where("user_id = ?", userID).
		Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	todoID := c.Param("id")
	var todo models.Todo

	if err := config.DB.First(&todo, todoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	userResp, err := clients.GetUserByID(todo.UserID)
	authorName := "Неизвестно"
	
	if err == nil {
		authorName = userResp.Username
	}

	c.JSON(http.StatusOK, gin.H{
		"todo":   todo,
		"author": authorName,
	})
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetUint("user_id")
    var todo models.Todo

    if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save changes"})
        return
    }

    config.DB.Preload("Category").Preload("Tags").First(&todo, todo.ID)

    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	result := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
