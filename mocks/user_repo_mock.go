package mocks

import (
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByCredentials(username, password string) (*models.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}
func (m *MockUserRepository) FindUserById(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}
