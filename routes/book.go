package routes

import (
	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getBooks",
	})
}
func getBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getBook",
	})

}
func createBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createBook",
	})

}
func updateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateBook",
	})

}
func deleteBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteBook",
	})
}
