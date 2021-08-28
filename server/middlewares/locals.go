package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func SetLocals(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		panic(err)
	}

	ID := session.Get("User")

	if ID != nil {
		u := models.User{}
		if err := u.PopulateFrom(fmt.Sprintf("SELECT * FROM users WHERE ID = '%d'", ID)); err != nil {
			return c.Redirect("/")
		}

		data := map[string]interface{}{
			"ID":        u.ID,
			"Email":     u.Email,
			"AuthType":  u.AuthType,
			"CreatedAt": u.CreatedAt,
			"UpdatedAt": u.UpdatedAt,
			"Admin":     u.Admin,
		}

		c.Locals("User", data)
	} else {
		c.Locals("User", nil)
	}

	c.Locals("Alert", session.Get("Alert"))

	session.Delete("Alert")
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.Next()
}
