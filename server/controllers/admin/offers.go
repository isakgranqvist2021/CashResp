package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetOffers(c *fiber.Ctx) error {
	return c.Render("pages/admin/offers", fiber.Map{
		"Title": "Offers",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}

func PostOffers(c *fiber.Ctx) error {
	var offer models.Offer

	if err := c.BodyParser(&offer); err != nil {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "could not parse body",
		})
	}

	if err := offer.CreateNew(); err != nil {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "error occured while creating offer",
		})
	}

	return c.Redirect(c.OriginalURL())
}
