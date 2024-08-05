package database

import (
	"ginserver/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dataSourceName string) {
	var err error

  var dbLogger logger.Interface

  if os.Getenv("GIN_MODE") != "debug" {
    dbLogger = logger.Discard
  }

	DB, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{
    Logger: dbLogger,
  })

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create a table with the User structure
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the users table  %v", err)
	}
}
