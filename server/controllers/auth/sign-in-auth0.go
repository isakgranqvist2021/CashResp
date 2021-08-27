package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/utils"
)

func SignInAuth0(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	state := base64.StdEncoding.EncodeToString(b)

	session.Set("state", state)
	err = session.Save()

	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	authenticator, err := utils.NewAuthenticator()
	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	return c.Redirect(authenticator.Config.AuthCodeURL(state))
}
