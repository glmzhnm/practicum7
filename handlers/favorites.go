package handlers

import (
	"awesomeProject3/database"
	"awesomeProject3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavoriteBooks(c *gin.Context) {
	uid, _ := c.Get("userID")
	userID := uid.(int)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	var favBooks []models.Book

	database.DB.Table("books").
		Joins("JOIN favorite_books ON favorite_books.book_id = books.id").
		Where("favorite_books.user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&favBooks)

	c.JSON(http.StatusOK, gin.H{
		"data":  favBooks,
		"page":  page,
		"limit": limit,
	})
}

func AddToFavorites(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	uid, _ := c.Get("userID")
	userID := uid.(int)

	fav := models.FavoriteBook{
		UserID: uint(userID),
		BookID: uint(bookID),
	}

	if err := database.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not add to favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added to favorites"})
}

func RemoveFromFavorites(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	uid, _ := c.Get("userID")
	userID := uid.(int)

	database.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.FavoriteBook{})
	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites"})
}
