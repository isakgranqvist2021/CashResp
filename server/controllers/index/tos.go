package index

import "github.com/gofiber/fiber/v2"

func GetTOS(c *fiber.Ctx) error {
	return c.Render("pages/index/tos", fiber.Map{
		"Title": "Terms Of Service",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
