package db

import (
	"log"

	"github.com/Studio-Centaurus/SlamjamAPI/config"
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.AppConfig.JWTSecret

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database:", err)
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln("failed to migrate database")
	}

	log.Println("Database connection and migration successful")
}