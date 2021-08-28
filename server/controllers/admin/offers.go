package admin

import "github.com/gofiber/fiber/v2"

func GetOffers(c *fiber.Ctx) error {
	return c.Render("pages/admin/offers", fiber.Map{
		"Title": "Offers",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}

func PostOffers(c *fiber.Ctx) error {
	return c.Redirect(c.OriginalURL())
}
