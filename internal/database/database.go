package database

import (
	"fiber-user-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DB struct {
	Db *gorm.DB
}

var Database DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Connected to database")

	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: add migrations
	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	Database.Db = db
}
