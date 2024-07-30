package routes

import (
	"github.com/Studio-Centaurus/SlamjamAPI/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App, userController *controller.UserController) {

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/user/signup", userController.Signup)
	app.Post("/user/login", userController.Login)
}
