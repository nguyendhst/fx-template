package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name             string               `json:"name"`
	ID               string               `json:"id"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}

type JwtCustomRefreshClaims struct {
	ID               string               `json:"id"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}

func (j *JwtCustomClaims) Valid() error {
	return j.RegisteredClaims.Valid()
}

func (j *JwtCustomRefreshClaims) Valid() error {
	return j.RegisteredClaims.Valid()
}
