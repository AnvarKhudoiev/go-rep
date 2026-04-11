package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookIDParam := c.Param("bookId")
	bookID, err := strconv.Atoi(bookIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var fav models.FavoriteBook
	err = config.DB.Where("user_id = ? AND book_id = ?", userID, bookID).First(&fav).Error

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already in favorites"})
		return
	}

	newFav := models.FavoriteBook{
		UserID: userID,
		BookID: uint(bookID),
	}

	if err := config.DB.Create(&newFav).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(200, newFav)

}


func RemoveFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookIDParam := c.Param("bookId")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	result := config.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Delete(&models.FavoriteBook{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed"})
}


func GetFavoriteBooks(c *gin.Context) {
	userID := c.GetUint("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var books []models.Book

	if err := config.DB.
		Joins("JOIN favorite_books ON favorite_books.book_id = books.id").
		Where("favorite_books.user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&books).Error; err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}