package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/models"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("C:/Users/yigit/Playground/opin.io/server/OpinionDatabase.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	//TODO: Add migrations.
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		return
	}

	Database = DBInstance{DB: db}

}
