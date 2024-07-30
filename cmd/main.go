package main

import (
	"github.com/Studio-Centaurus/SlamjamAPI/config"
	"github.com/Studio-Centaurus/SlamjamAPI/controller"
	db "github.com/Studio-Centaurus/SlamjamAPI/database"
	_ "github.com/Studio-Centaurus/SlamjamAPI/docs"
	"github.com/Studio-Centaurus/SlamjamAPI/repos"
	"github.com/Studio-Centaurus/SlamjamAPI/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()

	db.Connect()

	userRepo := &repos.UserRepository{DB: db.DB}	
	
	userController := &controller.UserController{Repo: userRepo}
	
	app := fiber.New()

	routes.Setup(app, userController)

	app.Listen(":8000")
}
