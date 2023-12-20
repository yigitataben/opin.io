package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/config"
	"server/models"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	configData, err := config.LoadConfigData()
	if err != nil {
		log.Fatal("Failed to load config data! \n", err.Error())
		os.Exit(2)
	}

	db, err := gorm.Open(sqlite.Open(configData.Database.FilePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	// TODO: Add migrations.
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})
	if err != nil {
		log.Fatal("Failed to run migrations! \n", err.Error())
		os.Exit(2)
	}

	Database = DBInstance{DB: db}
}
