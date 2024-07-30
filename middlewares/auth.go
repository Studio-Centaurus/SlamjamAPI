package middlewares

import (
	"github.com/Studio-Centaurus/SlamjamAPI/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
	})
}
