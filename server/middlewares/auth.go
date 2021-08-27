package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

func LoggedIn(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/auth/sign-in")
	}

	if session.Get("User") != nil {
		return c.Next()
	}

	return c.Redirect("/auth/sign-in")
}

func LoggedOut(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return c.Redirect("/account/profile")
	}

	if session.Get("User") == nil {
		return c.Next()
	}

	return c.Redirect("/users/profile")

}
