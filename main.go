package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/books", getBooks)
	// router.GET("/books/{id}",getBook)
	// router.POST("/books",createBook)
	// router.PUT("/books/{id}",updateBook)
	// router.DELETE("/books/{id}",deleteBook)

	router.Run(":8000")

}
