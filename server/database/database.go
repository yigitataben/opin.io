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
	db, err := gorm.Open(sqlite.Open("Opinion Database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed toconnect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	//TODO: Add migrations.
	db.AutoMigrate(&models.User{}, &models.Post{})

	Database = DBInstance{DB: db}

}
