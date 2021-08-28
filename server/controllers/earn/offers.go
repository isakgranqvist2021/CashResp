package earn

import (
	"github.com/gofiber/fiber/v2"
)

func GetOffers(c *fiber.Ctx) error {
	return c.Render("pages/earn/offers", fiber.Map{
		"Title": "Offers",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
