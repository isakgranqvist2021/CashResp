package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

func SetLocals(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		panic(err)
	}

	c.Locals("User", session.Get("User"))
	c.Locals("Alert", session.Get("Alert"))

	session.Delete("Alert")
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.Next()
}
