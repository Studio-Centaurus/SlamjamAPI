package repos

import (
	"errors"
	"log"

	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/Studio-Centaurus/SlamjamAPI/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) CreateUser(user models.User) error {
	if r.DB == nil {
		log.Println("database connection is not initialized")
		return errors.New("database connection is not initialized")
	}

	log.Printf("Creating user: %+v\n", user)
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing password")
		return err
	}
	user.Password = hashedPassword

	if err := r.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	log.Println("User created successfully")
	return nil
}

func (r *UserRepositoryImpl) FindByCredentials(username, password string) (*models.User, error) {
	var user models.User
	res := r.DB.Where("user_name = ?", username).First(&user)
	if res.Error != nil {
		log.Println("user not found")
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password incorrect")
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	var user models.User

	res := r.DB.Where("user_name = ?", username).First(&user)
	if res.Error != nil {
		log.Println("username not found")
		return nil, errors.New("user not found")
	}
	return &user, nil
}
func (r *UserRepositoryImpl) FindUserById(id string) (*models.User, error) {
	var user models.User

	res := r.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		log.Println("id anot found")
		return nil, errors.New("user not found")
	}
	return &user, nil
}
