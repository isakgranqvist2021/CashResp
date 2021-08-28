package earn

import (
	"github.com/gofiber/fiber/v2"
)

func GetSurveys(c *fiber.Ctx) error {
	return c.Render("pages/earn/surveys", fiber.Map{
		"Title": "Surveys",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}
