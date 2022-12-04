package database

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (db *gorm.DB) {
	var err error
	// Create database connection
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	// Set db log level
	db.Logger = db.Logger.LogMode(logger.Info)

	return
}
