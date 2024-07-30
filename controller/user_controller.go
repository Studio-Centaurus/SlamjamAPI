package controller

import (
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/repos"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo *repos.UserRepository
}

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
