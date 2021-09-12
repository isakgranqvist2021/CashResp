package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/routers"
	"github.com/isakgranqvist2021/cashresp/types"
)

func main() {
	app := fiber.New()

	gateway := app.Group("/api")
	routers.Users(gateway.Group("/users"))
	routers.Auth(gateway.Group("/auth"))

	app.All("*", func(c *fiber.Ctx) error {
		return c.JSON(types.Response{
			Message:    "not found",
			Success:    false,
			StatusCode: 404,
			Data:       nil,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

/*


 */
