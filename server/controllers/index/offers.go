package index

import (
	"github.com/gofiber/fiber/v2"
)

func GetOffers(c *fiber.Ctx) error {
	return c.Render("pages/index/offers", fiber.Map{
		"Title": "Offers",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
