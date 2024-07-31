package controller

import (
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/repos"
	"github.com/Studio-Centaurus/SlamjamAPI/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo repos.UserRepository
}

// Signup godoc
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
	NewUser, err := c.Repo.FindByCredentials(user.UserName, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	token, err := utils.CreateJwtToken(*NewUser)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user, token)
}

// Login godoc
// @Summary Login a user
// @Tags user
// @Accept mpfd
// @Produce json
// @Param loginRequest body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse
// @Router /user/login [post]
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
	token, err := utils.CreateJwtToken(*user)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Token":   token,
		"Message": "user has logged in",
	})
}

// GetUser godoc
// @Summary Get a user
// @Tags user
// @Accept mpfd
// @Produce json
// @Param UserRequest body models.UserRequest true "User Request"
// @Success 200 {object} models.User
// @Router /user/GetUser [post]
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	userRequest := new(models.UserRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := c.Repo.FindUserByNameAndID(userRequest.Username, userRequest.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})

}
