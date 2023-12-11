package router

import (
	"os"
	"time"

	"github.com/HsiaoCz/foo/api/av1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Router(addr string) error {
	r := fiber.New()
	file, err := os.OpenFile("../foo.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	r.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       file,
	}))
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.Post("/singup", av1.UserSingup)
				user.Post("/login", av1.UserLogin)
			}
		}
	}
	return r.Listen(addr)
}
