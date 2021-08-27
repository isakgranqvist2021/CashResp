package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetSignOut(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		fmt.Println(err)
	}

	if err := session.Destroy(); err != nil {
		fmt.Println(err)
	}

	return c.Redirect("/auth/sign-in")
}
