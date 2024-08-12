package repos

import "github.com/Studio-Centaurus/SlamjamAPI/models"

type TeamRepository interface {
	CreateTeam(team models.Team) error
	AddMembers(team *models.Team, users []models.User) error
	GetUsername(usernames []string, users *[]models.User) error
	GetUserById(ids []uint, users *[]models.User) error
	FindTeamByName(id string, team *models.Team) error
	// getTeamMembers(teamName string) (*[]models.User, error)
	// findTeamByName(teamName string) (*models.Team, error)
}
