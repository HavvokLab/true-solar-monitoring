package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserContext struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string     `json:"access_token"`
	ExpiredAt   *time.Time `json:"expired_at"`
}

type AccessToken struct {
	DisplayName string `json:"display_name"`
	jwt.RegisteredClaims
}

func (t *AccessToken) GetDisplayName() string {
	return t.DisplayName
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
