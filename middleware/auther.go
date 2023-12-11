package middleware

import "github.com/gofiber/fiber/v2"

func JWTAuther(c *fiber.Ctx) error {
	return c.Next()
}
