package models

import "github.com/dgrijalva/jwt-go"

//Claim Estructura con el token
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
