package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// RoleMiddleware checks if user has the required role
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ดึงข้อมูล JWT token จาก context
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// ตรวจสอบ role
		if claims["role"] != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: insufficient permissions",
			})
		}

		return c.Next()
	}
}
