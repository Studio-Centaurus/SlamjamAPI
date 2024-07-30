package utils

import (
	"time"

	"github.com/Studio-Centaurus/SlamjamAPI/config"
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CreateJwtToken(user models.User, c *fiber.Ctx) error {
	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(day * 1).Unix(),
		"issu":     time.Now(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(models.LoginResponse{
		Token: t,
	})

}
