package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
	"github.com/joho/godotenv"
)

func GetSignUp(c *fiber.Ctx) error {
	return c.Render("pages/auth/sign-up", fiber.Map{
		"Title":       "Sign Up",
		"User":        c.Locals("User"),
		"Alert":       c.Locals("Alert"),
		"Stylesheets": []string{"auth.min.css"},
		"Data":        nil,
	})
}

func PostSignUp(c *fiber.Ctx) error {
	u := models.User{
		AuthType:      "form",
		EmailVerified: false,
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := c.BodyParser(&u); err != nil {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  "an unexpected error has occured",
		})
	}

	if !utils.LongEnough(u.Email, 5, 499) {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  "email too short",
		})
	}

	if !utils.LongEnough(u.Password, 12, 499) {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  "password too short",
		})
	}

	if err := u.Register(); err != nil {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	if err := u.SetVerifyEmailAndSend(); err != nil {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  "mail could not be sent",
		})
	}

	return controllers.RedirectWithAlert(c, "/", utils.Alert{
		Severity: "success",
		Message:  fmt.Sprintf("an email has been sent to %s with instructions on how to verify your account", u.Email),
	})
}
