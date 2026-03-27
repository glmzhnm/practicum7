package handlers

import (
	"awesomeProject3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{}
var nextCategoryID = 1

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{ID: nextCategoryID, Name: input.Name}
	categories = append(categories, category)
	nextCategoryID++

	c.JSON(http.StatusCreated, category)
}
