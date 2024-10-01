package model

import "github.com/golang-jwt/jwt"

type JwtCustomClams struct {
	Userid   string
	RoleUser string
	jwt.StandardClaims
}
