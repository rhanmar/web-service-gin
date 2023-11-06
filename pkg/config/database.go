package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *sqlx.DB

func connectSqlite() (db *gorm.DB, err error) {
	database, err := gorm.Open(sqlite.Open("my_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database, err
}

//func connectPostgresGORM() (db *gorm.DB, err error) {
//	dsn := "postgres://postgres:postgres@localhost:15222"
//	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect to database!")
//	}
//	return database, err
//}

func connectPostgres() (db *sqlx.DB, err error) {
	dsn := "postgres://postgres:postgres@localhost:15222?sslmode=disable"
	database, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err2 := database.Ping()
	if err2 != nil {
		panic(err2)
	}

	return database, err
}

func ConnectDatabase() (*sqlx.DB, error) {
	//database, _ := connectSqlite()
	database, err := connectPostgres()
	if err != nil {
		return nil, err
	}
	//database.AutoMigrate(&models.Album{})
	DB = database

	return database, nil
}

//func ConnectTestDatabase() {
//	database, err := gorm.Open(sqlite.Open("my_test_db.db"), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect to database!")
//	}
//	database.AutoMigrate(&models.Album{})
//	DB = database
//}
//
//func ClearTestDatabase() {
//	DB.Where("id > 0").Delete(&models.Album{})
//}
