package services

import (
	"samples/spotify-proxy/models"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type Context struct {
	Properties     map[string]string
	Config         *models.Config
	SpotifyContext *SpotifyContext
}

func NewContext(config models.Config) *Context {
	return &Context{
		Properties: make(map[string]string),
		Config:     &config,
	}
}

type SpotifyContext struct {
	config *clientcredentials.Config
	client *spotify.Client
}

func NewSpotifyContext(config clientcredentials.Config, client spotify.Client) *SpotifyContext {
	return &SpotifyContext{
		config: &config,
		client: &client,
	}
}
