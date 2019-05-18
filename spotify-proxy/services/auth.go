package services

import (
	"context"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type AuthService interface {
	AuthenticateClient(ctx *Context) (*SpotifyContext, error)
}

type authService struct {
	logger Logger
}

func NewAuthService(logger Logger) AuthService {
	return &authService{
		logger: logger,
	}
}

func (svc *authService) AuthenticateClient(ctx *Context) (*SpotifyContext, error) {
	config := clientcredentials.Config{
		ClientID:     ctx.Config.ClientID,
		ClientSecret: ctx.Config.Secret,
		TokenURL:     ctx.Config.TokenEndpoint,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return nil, err
	}
	client := spotify.Authenticator{}.NewClient(token)
	return NewSpotifyContext(config, client), nil
}
