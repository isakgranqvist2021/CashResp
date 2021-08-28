package index

import "github.com/gofiber/fiber/v2"

func GetPrivacy(c *fiber.Ctx) error {
	return c.Render("pages/index/privacy", fiber.Map{
		"Title": "Your Privacy",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
