package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetClientCreateRoutes Creacion de usuario
func SetClientCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var client models.Client
	err := c.Bind(&client)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego cliente ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.ClientCreate(client, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientGetRoutes Creacion de usuario
func SetClientGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cl models.Client
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	cl.ID = id
	controllers.ClientGet(cl, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientGetListRoutes traer tipo de negocio
func SetClientGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var td models.Client
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de tipo de negocio no valido"
		return commons.DisplayMessage(c, &m)
	}
	td.CodCollection = id
	controllers.ClientGetList(td, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientEditRoutes Creacion de usuario
func SetClientEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cl models.Client
	err := c.Bind(&cl)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	cl.ID = id
	controllers.ClientUpdate(cl, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientDeleteRoutes Creacion de usuario
func SetClientDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cl models.Client
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	cl.ID = id
	controllers.ClientDelete(cl, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· telefono
································································
······························································*/

//SetClientTelCreateRoutes Creacion de telefono cliente
func SetClientTelCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ct models.ClientTel
	err := c.Bind(&ct)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego telefono cliente ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.ClientTelCreate(ct, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientTelGetRoutes traer telefono cliente
func SetClientTelGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ct models.ClientTel
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	ct.ID = id
	controllers.ClientTelGet(ct, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientTelGetListRoutes listado de telefono cliente
func SetClientTelGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ct models.ClientTel
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	ct.CodClient = id
	controllers.ClientTelGetList(ct, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientTelEditRoutes actualizar telefono cliente
func SetClientTelEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ct models.ClientTel
	err := c.Bind(&ct)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego telefono cliente ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	ct.ID = id
	controllers.ClientTelUpdate(ct, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientTelDeleteRoutes borrar telefono cliente
func SetClientTelDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ct models.ClientTel
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	ct.ID = id
	controllers.ClientTelDelete(ct, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· ubicaciones
································································
······························································*/

//SetClientListLocationCreateRoutes Creacion de Ubicacion
func SetClientListLocationCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cll models.ClientListLocation
	err := c.Bind(&cll)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego Ubicacion ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.ClientListLocationCreate(cll, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientListLocationGetRoutes traer Ubicacion
func SetClientListLocationGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cll models.ClientListLocation
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Ubicacion no valido"
		return commons.DisplayMessage(c, &m)
	}
	cll.ID = id
	controllers.ClientListLocationGet(cll, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientListLocationGetListRoutes listado de ubicaiones
func SetClientListLocationGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cll models.ClientListLocation
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de ubicacion no valido"
		return commons.DisplayMessage(c, &m)
	}
	cll.CodCollection = id
	controllers.ClientListLocationGetList(cll, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientListLocationEditRoutes actualizar Ubicacion
func SetClientListLocationEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cll models.ClientListLocation
	err := c.Bind(&cll)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego Ubicacion ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Ubicacion no valido"
		return commons.DisplayMessage(c, &m)
	}
	cll.ID = id
	controllers.ClientListLocationUpdate(cll, &m)
	return commons.DisplayMessage(c, &m)
}

//SetClientListLocationDeleteRoutes borrar Ubicacion
func SetClientListLocationDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cll models.ClientListLocation
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de Ubicacion no valido"
		return commons.DisplayMessage(c, &m)
	}
	cll.ID = id
	controllers.ClientListLocationDelete(cll, &m)
	return commons.DisplayMessage(c, &m)
}
