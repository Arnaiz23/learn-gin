package database

import (
	"ginserver/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create a table with the User structure
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the users table  %v", err)
	}
}
