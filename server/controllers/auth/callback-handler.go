package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func CallbackHandler(c *fiber.Ctx) error {
	session, err := utils.Store.Get(c)

	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	if c.Query("state") != session.Get("state") {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "Invalid state parameter",
			Message:  err.Error(),
		})

	}

	authenticator, err := utils.NewAuthenticator()
	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	token, err := authenticator.Config.Exchange(context.TODO(), c.Query("code"))
	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	oidcConfig := &oidc.Config{
		ClientID: "bkTpNm49zqJOlqepTqzTotGOLlSPSYJG",
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	u := models.User{
		AuthType:      profile["sub"].(string),
		Email:         profile["nickname"].(string),
		Password:      utils.RandKey(100, false),
		EmailVerified: true,
	}

	err = u.Register()

	if err != nil && err.Error() != "1" {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  "internal server error",
		})
	}

	if err != nil && err.Error() == "1" {
		u.PopulateFrom(fmt.Sprintf("SELECT * FROM users WHERE email = %s", u.Email))
	}

	payload := map[string]interface{}{
		"Profile": map[string]interface{}{
			"ID":       u.ID,
			"Email":    u.Email,
			"AuthType": u.AuthType,
		},
	}

	session.Set("User", payload)
	if err = session.Save(); err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	fmt.Println("----------- new sign in -----------")
	fmt.Printf("%d | %s | %s \n", u.ID, u.Email, u.AuthType)

	return controllers.RedirectWithAlert(c, "/users/profile", utils.Alert{
		Severity: "success",
		Message:  "sign in successful",
	})
}
