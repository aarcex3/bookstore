package routes

import (
	"net/http"

	"github.com/aarcex3/bookstore/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
func GetBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getBook",
	})

}
func CreateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createBook",
	})

}
func UpdateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateBook",
	})

}
func DeleteBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteBook",
	})
}
