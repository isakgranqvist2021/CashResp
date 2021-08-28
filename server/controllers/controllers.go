package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

type Injection struct {
	Title string
	Alert utils.Alert
	Data  interface{}
}

func RedirectWithAlert(c *fiber.Ctx, path string, alert utils.Alert) error {
	fmt.Println("alert -> ", alert)

	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/")
	}

	session.Set("Alert", alert)

	keys := session.Keys()
	fmt.Println(keys)

	if err := session.Save(); err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	return c.Redirect(path)
}
