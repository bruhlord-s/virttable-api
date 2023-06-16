package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}