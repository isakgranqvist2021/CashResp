package users

import "github.com/gofiber/fiber/v2"

func GetMyEarnings(c *fiber.Ctx) error {
	return c.Render("pages/users/earnings", fiber.Map{
		"Title": "My Earnings",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
