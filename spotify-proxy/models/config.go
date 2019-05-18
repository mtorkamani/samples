package models

// Config contains all configuration data
type Config struct {
	Port          string //Application port
	ClientID      string //Spotify Client ID to authenticate requests
	Secret        string //Spotify Client Secret to authenticate requests
	TokenEndpoint string //Spotify OAuth2 Endpoint URL at which access token requests go
	AuthEndpoint  string //Spotify OAuth2 Endpoint URL at whcih authorization requests go
}
