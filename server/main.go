package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/surveys/middlewares"
	"github.com/isakgranqvist2021/surveys/routers"
	"github.com/isakgranqvist2021/surveys/utils"
	"github.com/joho/godotenv"
)

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	// read .env file in root directory
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	utils.CreateStore()
	utils.CreateTablesIfNotExists()

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", "./public")

	app.Use("*", middlewares.SetLocals)

	routers.Index(app.Group("/"))
	routers.Users(app.Group("/users", middlewares.LoggedIn))
	routers.Auth(app.Group("/auth", middlewares.LoggedOut))

	log.Fatal(app.Listen(":8080"))
}
