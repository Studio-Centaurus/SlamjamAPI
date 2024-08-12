package routes

import (
	"github.com/Studio-Centaurus/SlamjamAPI/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App, userController *controller.UserController, teamController *controller.TeamController) {

	// jwt := middlewares.NewAuthMiddleware(config.AppConfig.JWTSecret)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/user/signup", userController.Signup)
	app.Post("/user/login", userController.Login)
	app.Get("/user/username/:username", userController.GetUsername)
	app.Get("/user/id/:id", userController.GetId)

	app.Post("/team/create", teamController.CreateTeam)
	app.Post("/team/addMembers/:teamid", teamController.AddMembers)
	app.Get("/team/:teamid", teamController.GetTeam)
}
