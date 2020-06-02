package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//ValidateJWT se valida que el token que llega sea valido
func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		var m models.Message
		tokenString, err := getToken(c.Request())
		if err != nil {
			m.Code = http.StatusBadRequest
			m.Message = err.Error()
			return commons.DisplayMessage(c, &m)
		}

		verifyFuction := func(token *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		}
		token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, verifyFuction)
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				mErr := ""
				vErr := err.(*jwt.ValidationError)
				//se puede agregar un modo debug e imprmiei el eror encaos de que ocurra
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					mErr = "su token a expirado"
				default:
					mErr = "erro de validacion de token"
				}
				m.Code = http.StatusBadRequest
				m.Message = mErr
			default:
				m.Code = http.StatusBadRequest
				m.Message = "erro y punto"
			}
			return commons.DisplayMessage(c, &m)
		}
		if !token.Valid {
			m.Code = http.StatusUnauthorized
			m.Message = "Su token no es valido"
			return commons.DisplayMessage(c, &m)
		}

		user := token.Claims.(*models.Claim).User
		c.Set("user", user)
		return next(c)
	}
}

//getToken trae el token de params o de
func getToken(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		ah = r.URL.Query().Get("authorization")
		if ah == "" {
			return "", errors.New("no llego ninguna authorization")
		}
		return ah, nil
	}
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	}
	return "", errors.New("el header no contiene la palabra bearer")
}
