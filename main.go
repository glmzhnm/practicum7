package main

import (
	"awesomeProject3/database"
	"awesomeProject3/handlers"
	"awesomeProject3/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.User{},
		&models.Author{},
		&models.Category{},
		&models.Book{},
		&models.FavoriteBook{},
	)
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)

	authGroup := r.Group("/")
	authGroup.Use(handlers.AuthMiddleware())
	{
		authGroup.POST("/books", handlers.CreateBook)
		authGroup.PUT("/books/:id", handlers.UpdateBook)
		authGroup.DELETE("/books/:id", handlers.DeleteBook)

		authGroup.GET("/books/favorites", handlers.GetFavoriteBooks)
		authGroup.PUT("/books/:id/favorites", handlers.AddToFavorites)
		authGroup.DELETE("/books/:id/favorites", handlers.RemoveFromFavorites)
	}

	r.Run(":8081")
}
