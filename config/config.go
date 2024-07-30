package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	DB_URl    string
}

var AppConfig Config

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	AppConfig = Config{
		JWTSecret: os.Getenv("JWTSecret"),
		DB_URl:    os.Getenv("DB_URL"),
	}

}
