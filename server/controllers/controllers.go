package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

type NewLogin struct {
	Ctx     *fiber.Ctx
	Session *session.Session
	User    *models.User
}

func RedirectWithAlert(c *fiber.Ctx, path string, alert utils.Alert) error {
	// fmt.Printf("Set alert -> %v\n", alert)

	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/")
	}

	session.Set("Alert", alert)
	if err := session.Save(); err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	return c.Redirect(path)
}

func CreateSession(n *NewLogin) error {
	n.Session.Set("User", n.User.ID)
	if err := n.Session.Save(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------- new sign in --------")
	fmt.Printf("%d | %s | %s \n", n.User.ID, n.User.Email, n.User.AuthType)

	return RedirectWithAlert(n.Ctx, "/users/profile", utils.Alert{
		Severity: "success",
		Message:  "sign in successful",
	})
}
