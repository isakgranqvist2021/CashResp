package routers

import "github.com/gofiber/fiber/v2"

func Index(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("qwt")
	})
}
