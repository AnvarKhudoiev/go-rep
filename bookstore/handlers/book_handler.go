package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"bookstore/models"
)

var Books []models.Book
var bookID = 1

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, Books)
}

func CreateBook(c *gin.Context) {
	var b models.Book
	if err := c.BindJSON(&b); err != nil || b.Title == "" || b.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	b.ID = bookID
	bookID++
	Books = append(Books, b)
	c.JSON(http.StatusCreated, b)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range Books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated models.Book
	if err := c.BindJSON(&updated); err != nil || updated.Title == "" || updated.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	for i, b := range Books {
		if b.ID == id {
			updated.ID = id
			Books[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range Books {
		if b.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}