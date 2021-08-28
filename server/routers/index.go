package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/index"
)

func Index(r fiber.Router) {
	r.Get("/offers", index.GetOffers)
	r.Get("/your-privacy", index.GetPrivacy)
	r.Get("/customer-support", index.GetSupport)
	r.Get("/terms-of-service", index.GetTOS)
	r.Get("/", index.GetHome)
}
