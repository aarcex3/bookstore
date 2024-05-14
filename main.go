package main

import (
	"github.com/aarcex3/bookstore/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", routes.GetBooks)
	router.GET("/books/:id", routes.GetBook)
	router.POST("/books", routes.CreateBook)
	router.PUT("/books/:id", routes.UpdateBook)
	router.DELETE("/books/:id", routes.DeleteBook)

	router.Run("localhost:8000")

}
