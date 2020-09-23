package db

import (
	"github.com/fedeya/test-fiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil

// Get singleton to the connection of database
func Get() *gorm.DB {
	if db != nil {
		return db
	}

	conn, err := gorm.Open(sqlite.Open("test-fiber.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(
		&models.Product{},
	)

	return db
}
