package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetDocumentTypeCreateRoutes Creacion de tipo de negocio
func SetDocumentTypeCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var dt models.DocumentType
	err := c.Bind(&dt)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.DocumentTypeCreate(dt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDocumentTypeGetRoutes traer tipo de negocio
func SetDocumentTypeGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var dt models.DocumentType
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	dt.ID = uint8(id)
	controllers.DocumentTypeGet(dt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDocumentTypeGetListRoutes traer tipo de negocio
func SetDocumentTypeGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var dt models.DocumentType
	controllers.DocumentTypeGetList(dt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDocumentTypeEditRoutes actualizar tipo de negocio
func SetDocumentTypeEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var dt models.DocumentType
	err := c.Bind(&dt)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de negocio ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	dt.ID = uint8(id)
	controllers.DocumentTypeUpdate(dt, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDocumentTypeDeleteRoutes borrar tipo de negocio
func SetDocumentTypeDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var dt models.DocumentType
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	dt.ID = uint8(id)
	controllers.DocumentTypeDelete(dt, &m)
	return commons.DisplayMessage(c, &m)
}
