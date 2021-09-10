package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/controllers"
	"github.com/isakgranqvist2021/cashresp/routers"
	"github.com/isakgranqvist2021/cashresp/utils"
)

func main() {
	app := fiber.New()

	if err := utils.InitQraphQL(); err != nil {
		log.Fatal("An error occured while initializing GraphQL")
	}

	gateway := app.Group("/api")
	routers.Index(gateway.Group("/"))
	routers.Users(gateway.Group("/users"))

	app.Post("/graphql", controllers.GraphQL)

	log.Fatal(app.Listen(":3000"))
}
