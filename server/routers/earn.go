package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/earn"
)

func Earn(r fiber.Router) {
	r.Get("/surveys", earn.GetSurveys)
	r.Get("/offers", earn.GetOffers)
	r.Get("/watch", earn.GetVideos)
}
