package services

import (
	"os"

	"samples/spotify-proxy/models"

	"github.com/zmb3/spotify"
)

type ConfigService interface {
	Get() (models.Config, error)
}

type configService struct {
	logger Logger
}

func NewConfigService(logger Logger) ConfigService {
	return &configService{
		logger: logger,
	}
}

func (svc *configService) Get() (models.Config, error) {
	svc.logger.Log("Starting to prepare configurations")
	port := get("5000", "PORT")
	clientID := get("__CLIENT_ID__", "CLIENT_ID")
	secret := get("__CLIENT_SECRET__", "CLIENT_SECRET")
	tokenEndpoint := get(spotify.TokenURL, "TOKEN_ENDPOINT")
	authEndpoint := get(spotify.AuthURL, "AUTH_ENDPOINT")
	svc.logger.Log("Finished preparing configurations")
	return models.Config{
		Port:          ":" + port,
		ClientID:      clientID,
		Secret:        secret,
		TokenEndpoint: tokenEndpoint,
		AuthEndpoint:  authEndpoint,
	}, nil
}

func get(defaultValue, envName string) string {
	result := defaultValue
	if len(os.Getenv(envName)) > 0 {
		result = os.Getenv(envName)
	}
	return result
}
