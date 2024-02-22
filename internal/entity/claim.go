package entity

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserGUID string `json:"user_guid"`
	jwt.StandardClaims
}
