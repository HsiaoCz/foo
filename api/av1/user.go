package av1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UserSingup(c *fiber.Ctx) error {
	users := new(UserS)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": err.Error(),
		})
	}
	if users.Password != users.RePassword {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "please check the password and the re_password",
		})
	}
	if !strings.Contains(users.Email, "@") {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Please type a correct email",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Code":    fiber.StatusOK,
		"Message": "Signup successed!",
	})
}

func UserLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Code":    fiber.StatusOK,
		"Message": "Login successed!",
	})
}
