package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetExpenseCreateRoutes Creacion de gasto
func SetExpenseCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var e models.Expense
	err := c.Bind(&e)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego gasto ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.ExpenseCreate(e, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseGetRoutes traer gasto
func SetExpenseGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var e models.Expense
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	e.ID = id
	controllers.ExpenseGet(e, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseGetListRoutes listado de ubicaiones
func SetExpenseGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var e models.Expense
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	e.CodCollection = id
	controllers.ExpenseGetList(e, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseEditRoutes actualizar gasto
func SetExpenseEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var e models.Expense
	err := c.Bind(&e)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego gasto ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	e.ID = id
	controllers.ExpenseUpdate(e, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseDeleteRoutes borrar gasto
func SetExpenseDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var e models.Expense
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	e.ID = id
	controllers.ExpenseDelete(e, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
···············  descripcion de gasto
································································
······························································*/

//SetExpenseDescripCreateRoutes Creacion de tipo de gasto
func SetExpenseDescripCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ed models.ExpenseDescrip
	err := c.Bind(&ed)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de gasto ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.ExpenseDescripCreate(ed, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseDescripGetRoutes traer tipo de gasto
func SetExpenseDescripGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ed models.ExpenseDescrip
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	ed.ID = id
	controllers.ExpenseDescripGet(ed, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseDescripGetListRoutes listado de ubicaiones
func SetExpenseDescripGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ed models.ExpenseDescrip
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	ed.CodCollection = id
	controllers.ExpenseDescripGetList(ed, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseDescripEditRoutes actualizar tipo de gasto
func SetExpenseDescripEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ed models.ExpenseDescrip
	err := c.Bind(&ed)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego tipo de gasto ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	ed.ID = id
	controllers.ExpenseDescripUpdate(ed, &m)
	return commons.DisplayMessage(c, &m)
}

//SetExpenseDescripDeleteRoutes borrar tipo de gasto
func SetExpenseDescripDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ed models.ExpenseDescrip
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de gasto no valido"
		return commons.DisplayMessage(c, &m)
	}
	ed.ID = id
	controllers.ExpenseDescripDelete(ed, &m)
	return commons.DisplayMessage(c, &m)
}
