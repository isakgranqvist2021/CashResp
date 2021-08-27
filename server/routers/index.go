package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/index"
)

func Index(r fiber.Router) {
	r.Get("/offers", index.GetOffers)
	r.Get("/", index.GetHome)
}
