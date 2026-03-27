package handlers

import (
	"awesomeProject3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authors []models.Author
var nextAuthorID = 1

type AuthorInput struct {
	Name string `json:"name" binding:"required"`
}

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var input AuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{ID: nextAuthorID, Name: input.Name}
	authors = append(authors, author)
	nextAuthorID++

	c.JSON(http.StatusCreated, author)
}
