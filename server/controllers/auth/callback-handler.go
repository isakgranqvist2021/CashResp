package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/controllers"
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

	payload := map[string]interface{}{
		"Profile": map[string]interface{}{
			"AuthType": profile["sub"].(string),
			"Email":    profile["nickname"].(string),
		},
		"AccessToken": token.AccessToken,
	}

	session.Set("User", payload)
	if err = session.Save(); err != nil {
		return controllers.RedirectWithAlert(c, "/auth/sign-in", utils.Alert{
			Severity: "error",
			Message:  err.Error(),
		})
	}

	fmt.Println("----------- new sign in -----------")
	for k, v := range profile {
		fmt.Printf("%s -> %s\n", k, v)
	}

	return c.Redirect("/auth/sign-in")
}
