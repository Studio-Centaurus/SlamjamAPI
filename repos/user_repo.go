package repos

import (
	"errors"
	"log"

	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}

	log.Printf("Creating user: %+v\n", user)

	if err := r.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	log.Println("User created successfully")
	return nil
}
