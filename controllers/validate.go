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
		tokenString, err := getTokenFromAuthorizationHeader(c.Request())
		if err != nil {
			tokenString, err = getTokenFromURLParams(c.Request())
			if err != nil {
				m.Code = 102
				return commons.DisplayMessage(c, &m)
			}
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
					mErr = "erro de validacon de token"
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

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	}
	return "", errors.New("el header no contiene la palabra Bearer")
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}
