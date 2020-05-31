package routes

import (
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetLoginRoutes inicio de seccion
func SetLoginRoutes(e echo.Context) error {
	m := models.Message{
		Code:    203,
		Message: "prueba",
	}
	return commons.DisplayMessage(e, &m)
}
