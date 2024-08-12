package repos

import (
	"errors"
	"log"

	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"gorm.io/gorm"
)

type TeamRepositoryImpl struct {
	DB *gorm.DB
}

func (r *TeamRepositoryImpl) CreateTeam(team models.Team) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	if err := r.DB.Create(&team).Error; err != nil {
		log.Println("error creating team")
		return err
	}
	log.Println("made team")
	return nil

}

func (r *TeamRepositoryImpl) FindTeamByName(id string, team *models.Team) error {
	return r.DB.Preload("Members").Where("id = ?", id).First(team).Error
}

func (r *TeamRepositoryImpl) AddMembers(team *models.Team, users []models.User) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	return r.DB.Model(team).Association("Members").Append(users)

}

func (r *TeamRepositoryImpl) GetUserById(ids []uint, users *[]models.User) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	return r.DB.Where("id IN ?", ids).Find(users).Error
}

func (r *TeamRepositoryImpl) GetUsername(usernames []string, users *[]models.User) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}
	return r.DB.Where("username IN ?", usernames).Find(users).Error
}
