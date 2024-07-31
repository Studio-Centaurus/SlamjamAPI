package repos

import "github.com/Studio-Centaurus/SlamjamAPI/models"

type UserRepository interface {
	CreateUser(user models.User) error
	FindByCredentials(username, password string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserById(id string) (*models.User, error)
}
