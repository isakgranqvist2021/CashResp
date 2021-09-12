package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/cashresp/controllers/auth"
)

func Auth(r fiber.Router) {
	r.Post("/login", auth.Login)
	r.Post("/register", auth.Register)
}
