package domain

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type AccessToken struct {
	DisplayName string `json:"display_name"`
	jwt.RegisteredClaims
}

func (t *AccessToken) GetDisplayName() string {
	return t.DisplayName
}
