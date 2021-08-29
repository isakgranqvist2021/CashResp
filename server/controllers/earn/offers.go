package earn

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetOffers(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/users/sign-out")
	}

	offers := []map[string]string{
		{
			"Href":      fmt.Sprintf("https://www.offertoro.com/ifr/show/29892/%d/13859", session.Get("User")),
			"Thumbnail": "",
			"Label":     "Offer Toro",
		},
	}

	return c.Render("pages/earn/offers", fiber.Map{
		"Title": "Offers",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data": map[string]interface{}{
			"Offers": offers,
		},
	})
}
