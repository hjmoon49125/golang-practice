package model

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}