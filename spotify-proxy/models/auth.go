package models

// AuthToken contains received access token data
type AuthToken struct {
	AccessToken string `json:"access_token" example:"ACCESS_TOKEN"`
	TokenType   string `json:"token_type" example:"ACCESS_TOKEN"`
	ExpiresIns  string `json:"expires_in" example:"ACCESS_TOKEN"`
	Score       string `json:"score" example:"ACCESS_TOKEN"`
}
