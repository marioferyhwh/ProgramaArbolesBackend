package routes

import (
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
)

//InitRoutes en rutamiento de la API
func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/api/v1/"
	e.POST(prefixAPI+"login", controllers.Login)

}
