package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/types"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(types.Response{
		Message:    "Log In",
		Success:    true,
		StatusCode: 200,
		Data:       nil,
	})
}
