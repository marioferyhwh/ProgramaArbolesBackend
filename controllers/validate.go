package controllers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

type key string

const keyUser key = "user"

//ValidateToken se valida que el token que llega sea valido
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)
	if err != nil {
		m.Code = 102
		return
	}
	if !token.Valid {
		m.Code = http.StatusUnauthorized
		m.Message = "Su token no es valido"
		return
	}

	ctx := context.WithValue(r.Context(), keyUser, token.Claims.(*models.Claim).User) // nolint
	next(w, r.WithContext(ctx))
}
