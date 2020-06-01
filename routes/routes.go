package routes

import (
	"github.com/labstack/echo"
)

//InitRoutes en rutamiento de la API
func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/api/v1/"
	const prefixUser = "user"
	e.POST(prefixAPI+"login", SetLoginRoutes)
	e.POST(prefixAPI+prefixUser, SetUserCreateRoutes)

}
