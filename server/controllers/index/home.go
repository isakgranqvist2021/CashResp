package index

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetHome(c *fiber.Ctx) error {
	if utils.SessionActive(c) {
		return c.Redirect("/users/profile")
	} else {
		return c.Redirect("/auth/sign-in")
	}
}
