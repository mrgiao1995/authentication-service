package model

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	ID       string
	IsActive bool
	UserName string
	Role     string
	jwt.RegisteredClaims
}
