package earn

import "github.com/gofiber/fiber/v2"

func GetEarnings(c *fiber.Ctx) error {
	return c.Render("pages/earn/earnings", fiber.Map{
		"Title": "My Earnings",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
