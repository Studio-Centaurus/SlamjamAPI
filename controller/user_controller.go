package controller

import (
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/repos"
	"github.com/Studio-Centaurus/SlamjamAPI/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo *repos.UserRepository
}

// Sighup godoc
// @Summary sighup a new user
// @Tages user
// @Accept mpfd
// @produce json
// @Success 200 {array} models.User
// @Router /user/signup [post]
func (c *UserController) Signup(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := c.Repo.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := c.Repo.FindByCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return utils.CreateJwtToken(*user, ctx)

}
