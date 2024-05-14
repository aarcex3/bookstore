package models

import (
	"github.com/aarcex3/bookstore/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	database.Connect()
	DB = database.GetDB()
	DB.AutoMigrate(&Book{})
}
