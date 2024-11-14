package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserID int    `json:"user_id"` // Perbaikan tag JSON
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
