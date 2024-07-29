package main

import (
	"github.com/Studio-Centaurus/SlamjamAPI/config"
	db "github.com/Studio-Centaurus/SlamjamAPI/database"
	_ "github.com/Studio-Centaurus/SlamjamAPI/docs"
	"github.com/Studio-Centaurus/SlamjamAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

    config.LoadEnv()

	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}