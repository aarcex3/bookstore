package routes

import (
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getBooks",
	})
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
