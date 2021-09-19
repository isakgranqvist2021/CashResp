package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/models"
	"github.com/isakgranqvist2021/cashresp/types"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)

		return c.JSON(types.Response{
			Message:    "error while parsing request body",
			Success:    false,
			StatusCode: 400,
			Data:       nil,
		})
	}

	return c.JSON(types.Response{
		Message:    "Register",
		Success:    true,
		StatusCode: 200,
		Data:       user,
	})
}
