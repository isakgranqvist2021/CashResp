package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

type Injection struct {
	Title string
	Alert utils.Alert
	Data  interface{}
}

func RedirectWithAlert(c *fiber.Ctx, path string, alert utils.Alert) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/")
	}

	session.Set("Alert", alert)

	if err := session.Save(); err != nil {
		return c.Redirect("/")
	}

	return c.Redirect(path)
}
