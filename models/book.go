package models

import (
	"github.com/aarcex3/bookstore/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
}

func Init() {
	database.Connect()
	DB = database.GetDB()
	DB.AutoMigrate(&Book{})
}
