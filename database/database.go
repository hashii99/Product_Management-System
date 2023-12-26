package database

import (
	// "database/sql"
	"log"
	"os"
	"test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		// return
		os.Exit(2)
	}

	log.Println("Connected to the database succefully!!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// todo - add migrations

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = Dbinstance{Db: db}



}
