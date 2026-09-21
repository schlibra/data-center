package models

import "github.com/golang-jwt/jwt/v5"

type JwtToken struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Enable   int    `json:"enable"`
	Admin    int    `json:"admin"`
	TokenID  string `json:"token_id"`
	jwt.RegisteredClaims
}
