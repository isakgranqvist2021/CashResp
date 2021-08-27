package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers/auth"
)

func Auth(r fiber.Router) {
	r.Get("/verify-email", auth.GetVerifyEmail)
	r.Get("/verify-email/:id", auth.PostVerifyEmail)
	r.Get("/sign-in", auth.GetSignIn)
	r.Post("/sign-in", auth.PostSignIn)
	r.Get("/sign-up", auth.GetSignUp)
	r.Post("/sign-up", auth.PostSignUp)

	r.Get("/sign-in/auth0", auth.SignInAuth0)
	r.Get("/callback", auth.CallbackHandler)
}
