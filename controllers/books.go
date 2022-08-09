package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookDTO struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

var DB map[string]BookDTO = make(map[string]BookDTO)

func BooksCreate(c *gin.Context) {
	// Get data from request body
	var dto BookDTO

	c.Bind(&dto)
	dto.Id = uuid.New()

	// Create a new Book
	// TODO using GORM
	DB[dto.Id.String()] = dto

	// Return it
	c.JSON(201, &dto)
}

func BooksGetAll(c *gin.Context) {
	// Get all from Database
	// TODO using GORM
	list := []BookDTO{}
	for _, value := range DB {
		list = append(list, value)
	}

	// Return it
	c.JSON(200, &list)
}

func BooksGetOne(c *gin.Context) {
	// Get One Database
	// TODO using GORM
	response, exist := DB[c.Param("id")]
	if !exist {
		c.JSON(404, gin.H{"message": "Book Not Found"})
		return
	}

	// Return it
	c.JSON(200, &response)
}

func BooksUpdate(c *gin.Context) {
	// Check if exist
	id := c.Param("id")
	data, exist := DB[id]
	if !exist {
		c.JSON(404, gin.H{"message": "Book Not Found"})
		return
	}

	// Bind DTO
	var dto BookDTO
	c.Bind(&dto)

	// Update
	DB[id] = BookDTO{Id: data.Id, Title: dto.Title}

	// Return it
	c.JSON(200, DB[id])
}

func BooksDelete(c *gin.Context) {
	// Check if exist
	id := c.Param("id")
	_, exist := DB[id]
	if !exist {
		c.JSON(404, gin.H{"message": "Book Not Found"})
		return
	}

	// Delete
	delete(DB, id)

	// Return it
	c.JSON(204, nil)
}
