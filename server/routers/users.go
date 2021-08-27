package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/users"
)

func Users(r fiber.Router) {
	r.Get("/profile", users.GetProfile)
	r.Get("/conversion", users.GetConversion)
	r.Get("/sign-out", users.GetSignOut)
}
