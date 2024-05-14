package routes

import (
	"net/http"

	"github.com/aarcex3/bookstore/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title           string `json:"title" binding:"required"`
	Author          string `json:"author" binding:"required"`
	PublicationDate string `json:"publication_date" binding:"required"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
func GetBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author, PublicationDate: input.PublicationDate}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})

}
func UpdateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateBook",
	})

}
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
