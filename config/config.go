package config

import (
	"os"

	"github.com/zi-bot/simple-gin-rest/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Asset{})

	return database
}

func ConnectDatabaseTest() *gorm.DB {
	_, err := os.Stat("../testfiles")
	if os.IsNotExist(err) {
		os.Mkdir("../testfiles", 0755)
	}
	os.Remove("../testfiles/test.db")
	database, err := gorm.Open(sqlite.Open("../testfiles/test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Asset{})

	return database
}
