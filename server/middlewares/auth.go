package middlewares

import (
	"github.com/gofiber/fiber/v2"
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
