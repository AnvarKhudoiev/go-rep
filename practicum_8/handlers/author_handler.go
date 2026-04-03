package handlers

import (
	"net/http"
	"practicum_8/config"
	"practicum_8/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&author)
	c.JSON(http.StatusOK, author)
}

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	config.DB.Preload("Todos").Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := config.DB.Preload("Todos").First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func CreateTodoForAuthor(c *gin.Context) {
	authorID := c.Param("id")
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(authorID)
	todo.AuthorID = uint(id)
	config.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}