package repos

import "github.com/Studio-Centaurus/SlamjamAPI/models"

type TeamRepository interface {
	createTeam(team models.Team) error
	getTeamMembers(teamName string) (*[]models.User, error)
	findTeamByName(teamName string) (*models.Team, error)
}
