package repos

import (
	"errors"
	"log"

	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/utils"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassoword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("errror hasehing passwond")
	}
	user.Password = hashedPassoword

	if err := r.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	log.Println("User created successfully")
	return nil
}

func (r *UserRepository) FindByCredentials(username, password string) (*models.User, error) {

	var user models.User

	res := r.DB.Where("user_name= ?", username).First(&user)
	if res.Error != nil {
		log.Println("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("passwod incorect")
	}

	return &user, nil

}
