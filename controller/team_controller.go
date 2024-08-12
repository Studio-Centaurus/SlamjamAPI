package controller

import (
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/repos"
	"github.com/gofiber/fiber/v2"
)

type TeamController struct {
	Repo repos.TeamRepository
}

func (c *TeamController) CreateTeam(ctx *fiber.Ctx) error {
	var team models.Team
	if err := ctx.BodyParser(&team); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid requst body",
		})
	}
	if err := c.Repo.CreateTeam(team); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "could not create user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(team)
}

func (c *TeamController) AddMembers(ctx *fiber.Ctx) error {
	teamID := ctx.Params("teamID")
	var request struct {
		UserIDs   []uint   `json:"user_ids"`
		Usernames []string `json:"usernames"`
	}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid requst body",
		})
	}

	//find team
	var team models.Team
	if err := c.Repo.FindTeamByName(teamID, &team); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid requst body",
		})

	}

	//find users
	var users []models.User
	if len(request.UserIDs) > 0 {
		if err := c.Repo.GetUserById(request.UserIDs, &users); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid requst body",
			})
		}
	} else if len(request.Usernames) > 0 {
		if err := c.Repo.GetUsername(request.Usernames, &users); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "could not fetch users by usernames",
			})
		}
	}

	//add to team
	if err := c.Repo.AddMembers(&team, users); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not add members to the team",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "members added successfully",
	})
}

func (c *TeamController) GetTeam(ctx *fiber.Ctx) error {
	teamID := ctx.Params("teamID")

	var team models.Team
	if err := c.Repo.FindTeamByName(teamID, &team); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not fetch users by usernames",
		})
	}

	return ctx.JSON(fiber.Map{
		"team": team,
	})
}
