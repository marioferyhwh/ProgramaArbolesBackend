package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

/*······························································
································································
··············· cobroes
································································
······························································*/

//SetCollectionCreateRoutes Creacion de cobro
func SetCollectionCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var co models.Collection
	err := c.Bind(&co)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.CollectionCreate(co, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionGetRoutes traer cobro
func SetCollectionGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var co models.Collection
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	co.ID = id
	controllers.CollectionGet(co, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionGetListRoutes listado de ubicaiones
func SetCollectionGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var co models.Collection
	controllers.CollectionGetList(co, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionEditRoutes actualizar cobro
func SetCollectionEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var co models.Collection
	err := c.Bind(&co)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	co.ID = id
	controllers.CollectionUpdate(co, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionDeleteRoutes borrar cobro
func SetCollectionDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var co models.Collection
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	co.ID = id
	controllers.CollectionDelete(co, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· Movimentos Cobroes
································································
······························································*/

//SetCollectionCashCreateRoutes Creacion de Movimentos Cobro
func SetCollectionCashCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ch models.CollectionCash
	err := c.Bind(&ch)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego Movimentos Cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.CollectionCashCreate(ch, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionCashGetRoutes traer Movimentos Cobro
func SetCollectionCashGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ch models.CollectionCash
	id, err := getParams64(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Movimentos Cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	ch.ID = id
	controllers.CollectionCashGet(ch, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionCashGetListRoutes listado de ubicaiones
func SetCollectionCashGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ch models.CollectionCash
	id, err := getParams32(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Movimentos Cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	ch.CodCollection = id
	controllers.CollectionCashGetList(ch, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionCashEditRoutes actualizar Movimentos Cobro
func SetCollectionCashEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ch models.CollectionCash
	err := c.Bind(&ch)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego Movimentos Cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Movimentos Cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	ch.ID = id
	controllers.CollectionCashUpdate(ch, &m)
	return commons.DisplayMessage(c, &m)
}

//SetCollectionCashDeleteRoutes borrar Movimentos Cobro
func SetCollectionCashDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ch models.CollectionCash
	id, err := getParams64(models.GetParams{C: c})
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Movimentos Cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	ch.ID = id
	controllers.CollectionCashDelete(ch, &m)
	return commons.DisplayMessage(c, &m)
}
