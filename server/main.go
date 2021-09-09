package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/routers"
)

func main() {
	app := fiber.New()

	gateway := app.Group("/api")

	routers.Index(gateway.Group("/"))
	routers.Users(gateway.Group("/users"))

	log.Fatal(app.Listen(":3000"))
}
