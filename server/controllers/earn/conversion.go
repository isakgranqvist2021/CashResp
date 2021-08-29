package earn

import "github.com/gofiber/fiber/v2"

func GetConversion(c *fiber.Ctx) error {

	return c.Redirect("/users/my-earnings")
}
