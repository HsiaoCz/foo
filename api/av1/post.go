package av1

import "github.com/gofiber/fiber/v2"

// user login struct
type UserL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// user signup
type UserS struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Email      string `json:"email"`
}

// some struct
type Post struct {
	Message  string `json:"message"`
	Identity string `json:"identity"`
}

// some envirment class
type Hero struct {
	Ctx  *fiber.Ctx
	Post Post
}
