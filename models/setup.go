package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("my_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Album{})
	DB = database
}

func ConnectTestDatabase() {
	database, err := gorm.Open(sqlite.Open("my_test_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Album{})
	DB = database
}

func ClearTestDatabase() {
	DB.Where("id > 0").Delete(&Album{})
}
