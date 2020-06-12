package routes

import (
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetGeneralValuesGetRoutes traer tipo de negocio
func SetGeneralValuesGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	controllers.GeneralValuesGet(&m)
	return commons.DisplayMessage(c, &m)
}
