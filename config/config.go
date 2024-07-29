package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
}

var AppConfig Config


func LoadEnv() {

	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading env")
	}

	// AppConfig := Config{
	// 	JWTSecret: os.Getenv("JWTSECRET"),		
	// }

}