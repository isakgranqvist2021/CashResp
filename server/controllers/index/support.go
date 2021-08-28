package index

import "github.com/gofiber/fiber/v2"

func GetSupport(c *fiber.Ctx) error {
	return c.Render("pages/index/support", fiber.Map{
		"Title": "Customer Support",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
