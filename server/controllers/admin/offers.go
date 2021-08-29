package admin

import (
	"fmt"

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

/*
type Offer struct {
	ID          int
	OfferID     string
	PublisherID string
	AppID       string
	Description string
	CreatedAt   string
	UpdatedAt   string
	Href        string
	ImageID     int
	Provider    string
}
*/
func PostOffers(c *fiber.Ctx) error {
	var offer models.Offer

	if err := c.BodyParser(&offer); err != nil {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "could not parse body",
		})
	}

	fmt.Println(offer.Draft)

	if !utils.LongEnough(offer.OfferID, 3, 45) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "missing offer id",
		})
	}

	if !utils.LongEnough(offer.PublisherID, 3, 45) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "missing publisher id",
		})
	}

	if !utils.LongEnough(offer.AppID, 3, 45) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "missing app id",
		})
	}

	if !utils.LongEnough(offer.Description, 3, 999) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "miss description",
		})
	}

	if !utils.LongEnough(offer.Href, 3, 45) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "missing href",
		})
	}

	if !utils.LongEnough(offer.Provider, 3, 45) {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "missing provider",
		})
	}

	if err := offer.CreateNew(); err != nil {
		return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
			Severity: "danger",
			Message:  "error occured while creating offer",
		})
	}

	return controllers.RedirectWithAlert(c, "/admin/offers", utils.Alert{
		Severity: "success",
		Message:  "offer has been created",
	})
}
