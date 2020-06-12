package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetTelDescripCreateRoutes Creacion de tipo de negocio
func SetTelDescripCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.TelDescrip
	err := c.Bind(&td)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.TelDescripCreate(td, &m)
	return commons.DisplayMessage(c, &m)
}

//SetTelDescripGetRoutes traer tipo de negocio
func SetTelDescripGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.TelDescrip
	id, err := getParams8(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	td.ID = id
	controllers.TelDescripGet(td, &m)
	return commons.DisplayMessage(c, &m)
}

//SetTelDescripGetListRoutes traer tipo de negocio
func SetTelDescripGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.TelDescrip
	id, err := getParams8(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	td.ID = id
	controllers.TelDescripGetList(td, &m)
	return commons.DisplayMessage(c, &m)
}

//SetTelDescripEditRoutes actualizar tipo de negocio
func SetTelDescripEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.TelDescrip
	err := c.Bind(&td)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams8(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	td.ID = id
	controllers.TelDescripUpdate(td, &m)
	return commons.DisplayMessage(c, &m)
}

//SetTelDescripDeleteRoutes borrar tipo de negocio
func SetTelDescripDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.TelDescrip
	id, err := getParams8(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	td.ID = id
	controllers.TelDescripDelete(td, &m)
	return commons.DisplayMessage(c, &m)
}
