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
	dsn := config.AppConfig.DbUrl
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.League{}, &models.Match{}, &models.Team{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database connection and migration successful")
}
