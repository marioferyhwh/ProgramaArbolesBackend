package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetBusinessTypeCreateRoutes Creacion de tipo de negocio
func SetBusinessTypeCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var bt models.BusinessType
	err := c.Bind(&bt)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.BusinessTypeCreate(bt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetBusinessTypeGetRoutes traer tipo de negocio
func SetBusinessTypeGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var bt models.BusinessType
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	bt.ID = uint8(id)
	controllers.BusinessTypeGet(bt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetBusinessTypeGetListRoutes traer tipo de negocio
func SetBusinessTypeGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var bt models.BusinessType
	controllers.BusinessTypeGetList(bt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetBusinessTypeEditRoutes actualizar tipo de negocio
func SetBusinessTypeEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var bt models.BusinessType
	err := c.Bind(&bt)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	bt.ID = uint8(id)
	controllers.BusinessTypeUpdate(bt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetBusinessTypeDeleteRoutes borrar tipo de negocio
func SetBusinessTypeDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var bt models.BusinessType
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	bt.ID = uint8(id)
	controllers.BusinessTypeDelete(bt, &m)
	return commons.DisplayMessage(c, &m)
}
