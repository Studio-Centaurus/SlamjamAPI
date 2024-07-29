package main

import (
	"github.com/Studio-Centaurus/SlamjamAPI/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

    config.LoadEnv()
	
	app := fiber.New()

	app.Listen(":8000")
}