package config

import (
	"example/web-service-gin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectSqlite() (db *gorm.DB, err error) {
	database, err := gorm.Open(sqlite.Open("my_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database, err
}

func connectPostgres() (db *gorm.DB, err error) {
	dsn := "postgres://postgres:postgres@localhost:15222"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database, err
}

func ConnectDatabase() {
	//database, _ := connectSqlite()
	database, _ := connectPostgres()
	//database.AutoMigrate(&models.Album{})
	DB = database
}

func ConnectTestDatabase() {
	database, err := gorm.Open(sqlite.Open("my_test_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&models.Album{})
	DB = database
}

func ClearTestDatabase() {
	DB.Where("id > 0").Delete(&models.Album{})
}
