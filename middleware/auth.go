package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")), // ใช้ Secret Key จาก .env
		ErrorHandler: jwtErrorHandler,                     // Custom Error Handler
	})
}

// jwtErrorHandler handles unauthorized access
func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}
	return nil
}
