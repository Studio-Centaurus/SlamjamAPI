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

// @title SlamjamAPI
// @version 1.0
// @description api to serve user info
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {

	config.LoadEnv()

	db.Connect()

	userRepo := &repos.UserRepositoryImpl{DB: db.DB}
	teamRepo := &repos.TeamRepositoryImpl{DB: db.DB}
	userController := &controller.UserController{Repo: userRepo}
	teamController := &controller.TeamController{Repo: teamRepo}

	app := fiber.New()

	// jwt := middlewares.NewAuthMiddleware(config.Secret)

	routes.Setup(app, userController, teamController)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
