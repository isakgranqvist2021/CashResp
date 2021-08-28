package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func LoggedIn(c *fiber.Ctx) error {
	if utils.SessionActive(c) {
		return c.Next()
	} else {
		return c.Redirect("/auth/sign-in")
	}
}

func LoggedOut(c *fiber.Ctx) error {
	if !utils.SessionActive(c) {
		return c.Next()
	} else {
		return c.Redirect("/users/profile")
	}
}

func IsAdmin(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/users/sign-out")
	}

	ID := session.Get("User")
	u := models.User{}
	if ID != nil {
		if err := u.PopulateFrom(fmt.Sprintf("SELECT * FROM users WHERE ID = '%d'", ID)); err != nil {
			return c.Redirect("/")
		}
	} else {
		return c.Redirect("/")
	}

	if utils.SessionActive(c) && u.Admin {
		return c.Next()
	} else {
		return c.Redirect("/users/sign-out")
	}
}
