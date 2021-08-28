package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/admin"
)

func Admin(r fiber.Router) {
	r.Get("/images", admin.GetImages)
	r.Post("/images", admin.PostImages)
}
