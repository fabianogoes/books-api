package main

import (
	"github.com/fabianogoes/books-api/controllers"
	"github.com/fabianogoes/books-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadParameters()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Welcome Books API"}) })
	r.POST("/books", controllers.BooksCreate)
	r.GET("/books", controllers.BooksGetAll)
	r.GET("/books/:id", controllers.BooksGetOne)
	r.PUT("/books/:id", controllers.BooksUpdate)
	r.DELETE("/books/:id", controllers.BooksDelete)
	r.Run()
}
