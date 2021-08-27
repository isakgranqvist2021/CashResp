package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

type Body struct {
	Code string
}

func GetVerifyEmail(c *fiber.Ctx) error {
	return c.Render("pages/index/verify-email", fiber.Map{
		"Title": "Sign Up",
		"User":  c.Locals("User"),
		"Alert": c.Locals("Alert"),
		"Data":  nil,
	})
}

func PostVerifyEmail(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		fmt.Println(err)
		return controllers.RedirectWithAlert(c, "/", utils.Alert{
			Severity: "error",
			Message:  "an unexpected error has occured",
		})
	}

	u := models.User{
		VerifyCode: c.Params("id"),
	}

	if err := u.VerifyEmail(); err != nil {
		return controllers.RedirectWithAlert(c, "/", utils.Alert{
			Severity: "error",
			Message:  "an unexpected error has occured",
		})
	}

	session.Set("User", u.Email)
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}

	return controllers.RedirectWithAlert(c, "/", utils.Alert{
		Severity: "success",
		Message:  "your email has been verified",
	})
}
