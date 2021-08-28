package users

import "github.com/gofiber/fiber/v2"

func GetWithdrawal(c *fiber.Ctx) error {
	return c.Render("pages/users/withdrawal", fiber.Map{
		"Title": "Withdrawal",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
