package handlers

import (
	"awesomeProject3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}
var nextBookID = 1

type BookInput struct {
	Title      string  `json:"title" binding:"required"`
	AuthorID   int     `json:"author_id" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
}

func GetBooks(c *gin.Context) {
	var filteredBooks []models.Book

	categoryName := c.Query("category")
	if categoryName != "" {
		var catID int
		for _, cat := range categories {
			if cat.Name == categoryName {
				catID = cat.ID
				break
			}
		}

		if catID != 0 {
			for _, b := range books {
				if b.CategoryID == catID {
					filteredBooks = append(filteredBooks, b)
				}
			}
		}
	} else {
		filteredBooks = books
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	start := (page - 1) * limit
	end := start + limit
	total := len(filteredBooks)

	if start >= total {
		c.JSON(http.StatusOK, gin.H{"data": []models.Book{}, "total": total, "page": page})
		return
	}
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  filteredBooks[start:end],
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		ID:         nextBookID,
		Title:      input.Title,
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
		Price:      input.Price,
	}
	books = append(books, book)
	nextBookID++

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, b := range books {
		if b.ID == id {
			books[i].Title = input.Title
			books[i].AuthorID = input.AuthorID
			books[i].CategoryID = input.CategoryID
			books[i].Price = input.Price
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
