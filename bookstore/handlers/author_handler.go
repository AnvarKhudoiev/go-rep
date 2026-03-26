package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"bookstore/models"
)

var Authors []models.Author
var authorID = 1

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, Authors)
}

func CreateAuthor(c *gin.Context) {
	var a models.Author
	if err := c.BindJSON(&a); err != nil || a.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	a.ID = authorID
	authorID++
	Authors = append(Authors, a)
	c.JSON(http.StatusCreated, a)
}