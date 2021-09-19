package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/models"
	"github.com/isakgranqvist2021/cashresp/types"
)

func Login(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(types.Response{
			Message:    "error while parsing request body",
			Success:    false,
			StatusCode: 400,
			Data:       nil,
		})
	}

	return c.JSON(types.Response{
		Message:    "Log In",
		Success:    true,
		StatusCode: 200,
		Data:       nil,
	})
}
