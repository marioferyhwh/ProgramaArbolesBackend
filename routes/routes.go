package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//InitRoutes en rutamiento de la API
func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/api/v1/"
	const prefixUser = "user"
	e.GET(prefixAPI+"login", SetLoginRoutes)
	e.POST(prefixAPI+prefixUser, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixAPI+prefixUser+"/:id", SetGetUserRoutes, controllers.ValidateJWT)
	e.PUT(prefixAPI+prefixUser+"/:id", SetEditUserRoutes, controllers.ValidateJWT)
	e.DELETE(prefixAPI+prefixUser+"/:id", SetDeleteUserRoutes, controllers.ValidateJWT)
	e.GET("/*", SetDefaultRoutes)
	e.POST("/*", SetDefaultRoutes)
	e.DELETE("/*", SetDefaultRoutes)
	e.PUT("/*", SetDefaultRoutes)

}

//SetDefaultRoutes Creacion de usuario
func SetDefaultRoutes(c echo.Context) error {

	m := models.Message{}

	m.Code = http.StatusBadRequest
	m.Message = "no existe esa pagina"

	return commons.DisplayMessage(c, &m)
}
