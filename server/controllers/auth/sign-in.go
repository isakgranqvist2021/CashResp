package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetSignIn(c *fiber.Ctx) error {
	return c.Render("pages/auth/sign-in", fiber.Map{
		"Title":       "Sign In",
		"User":        c.Locals("User"),
		"Alert":       c.Locals("Alert"),
		"Stylesheets": []string{"auth.min.css"},
		"Data":        nil,
	})
}

func PostSignIn(c *fiber.Ctx) error {
	var u models.User
	session, err := utils.Store.Get(c)

	if err != nil {
		return controllers.RedirectWithAlert(c, "/", utils.Alert{
			Severity: "error",
			Message:  "session has exipred",
		})
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

	if err := u.Login(); err != nil {
		return controllers.RedirectWithAlert(c, c.OriginalURL(), utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	session.Set("User", u.ID)
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------- new sign in --------")
	fmt.Printf("%d | %s | %s \n", u.ID, u.Email, u.AuthType)

	return controllers.RedirectWithAlert(c, "/users/profile", utils.Alert{
		Severity: "success",
		Message:  "sign in successful",
	})
}
