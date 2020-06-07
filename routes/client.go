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

//SetClientRoutes Creacion de usuario
func SetClientRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var cl models.Client
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de cliente no valido"
		return commons.DisplayMessage(c, &m)
	}
	cl.ID = id
	controllers.ClientGet(cl, &m)
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
	id, err := getid(c)
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
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	cl.ID = id
	controllers.ClientDelete(cl, &m)
	return commons.DisplayMessage(c, &m)
}
