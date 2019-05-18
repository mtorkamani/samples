package cfg

import (
	"os"
)

type ConfigService interface {
	Get() (Config, error)
}

type configService struct {
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (svc *configService) Get() (Config, error) {
	port := get("5000", "PORT")
	return Config{
		Port: ":" + port,
	}, nil
}

func get(defaultValue, envName string) string {
	result := defaultValue
	if len(os.Getenv(envName)) > 0 {
		result = os.Getenv(envName)
	}
	return result
}
