package utils

import (
	"context"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://dustiastheguy.eu.auth0.com/")

	if err != nil {
		return nil, err
	}

	keys, err := GetKeys()

	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     keys["0auth"]["client_id"],
		ClientSecret: keys["0auth"]["client_secret"],
		RedirectURL:  "http://localhost:8080/auth/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
