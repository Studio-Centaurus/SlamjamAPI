package repos

import "github.com/Studio-Centaurus/SlamjamAPI/models"

type UserRepository interface {
	CreateUser(user models.User) error
	FindByCredentials(username, password string) (*models.User, error)
	FindUserByNameAndID(username, password string) (*models.User, error)
}
