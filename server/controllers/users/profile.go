package users

import (
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	return c.Render("pages/users/profile", fiber.Map{
		"Title": "Profile",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
