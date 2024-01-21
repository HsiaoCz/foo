package middleware

import "github.com/gofiber/fiber/v2"

// JWT
func JWTAuther(c *fiber.Ctx) error {
	return c.Next()
}
