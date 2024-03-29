package models

import "github.com/4kayDev/admitad-integration/internal/utils/jsoner"

type Authorization struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Language     string `json:"language"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Group        string `json:"group"`
}

func String(a *Authorization) string {
	return jsoner.Jsonify(a)
}
