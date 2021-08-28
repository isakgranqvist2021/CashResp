package earn

import (
	"github.com/gofiber/fiber/v2"
)

func GetVideos(c *fiber.Ctx) error {
	return c.Render("pages/earn/videos", fiber.Map{
		"Title": "Watch & Earn",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
