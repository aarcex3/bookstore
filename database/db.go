package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
