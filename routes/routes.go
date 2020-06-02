package routes

import (
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
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

}
