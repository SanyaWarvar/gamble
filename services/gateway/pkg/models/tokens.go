package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshToken struct {
	Token   string    `json:"token" db:"token"`
	UserId  uuid.UUID `json:"user_id" db:"user_id"`
	ExpDate time.Time `json:"exp_date" db:"exp_date"`
}

type AccessTokenClaims struct {
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
	RefreshId uuid.UUID `json:"refresh_id" db:"refresh_id"`
	jwt.RegisteredClaims
}

type RefreshInput struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
