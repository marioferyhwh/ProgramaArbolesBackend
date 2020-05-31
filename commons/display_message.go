package commons

import (
	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//DisplayMessage estructura de respuesta
func DisplayMessage(c echo.Context, m *models.Message) error {
	return c.JSON(m.Code, m)
}
