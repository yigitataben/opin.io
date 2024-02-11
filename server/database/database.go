package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"server/config"
	"server/models"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() error {
	configData, err := config.LoadConfigData()
	if err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open(configData.Database.FilePath), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Connected to the database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	// TODO: Add migrations.
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})
	if err != nil {
		return err
	}

	Database = DBInstance{DB: db}
	return nil
}
