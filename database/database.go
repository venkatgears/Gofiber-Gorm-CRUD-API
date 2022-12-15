package database

import (
	"log"
	"os"

	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func ConnectDb() {
	db, err := gorm.Open((sqlite.Open("api.db")), &gorm.Config{})
	if err != nil {
		log.Fatal("connection not established ,\n", err.Error())
		os.Exit(2)
	}

	log.Println("connected to db")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	db.AutoMigrate(&models.Order{}, &models.Product{}, &models.User{})

	Database = Dbinstance{Db: db}
	// fmt.Print(Database)
}
