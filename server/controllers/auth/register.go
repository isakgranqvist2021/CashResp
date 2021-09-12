package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/types"
)

func Register(c *fiber.Ctx) error {
	return c.JSON(types.Response{
		Message:    "Register",
		Success:    true,
		StatusCode: 200,
		Data:       nil,
	})
}
