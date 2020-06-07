package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

func getid(c echo.Context) (uint64, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(i64), nil
}

func getid32(c echo.Context) (uint32, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i64), nil
}

func getUserInterface(c echo.Context, u *models.User) {
	user := c.Get(commons.User)
	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, u)
}

//SetLoginRoutes inicio de seccion
func SetLoginRoutes(c echo.Context) error {
	var user models.User
	m := models.Message{}
	err := c.Bind(&user)
	if err != nil {
		m.Code = http.StatusNotFound
		m.Message = fmt.Sprint("no", err.Error())
		return commons.DisplayMessage(c, &m)
	}
	controllers.Login(user, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserCreateRoutes Creacion de usuario
func SetUserCreateRoutes(c echo.Context) error {
	user := models.User{}
	m := models.Message{}
	getUserInterface(c, &m.User)
	err := c.Bind(&user)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	if user.ConfirmPassword != user.Password {
		m.Code = http.StatusBadRequest
		m.Message = "contraseñas no coninciden"
		return commons.DisplayMessage(c, &m)
	}
	controllers.UserCreate(user, &m)
	return commons.DisplayMessage(c, &m)

}

//SetUserGetRoutes Creacion de usuario
func SetUserGetRoutes(c echo.Context) error {
	m := models.Message{}
	getUserInterface(c, &m.User)
	var u models.User
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	u.ID = id
	u.GetListUser = true
	u.GetUserTel = true
	controllers.UserGet(u, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserEditRoutes Creacion de usuario
func SetUserEditRoutes(c echo.Context) error {

	m := models.Message{}
	getUserInterface(c, &m.User)

	var u models.User
	err := c.Bind(&u)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	if id != u.ID {
		m.Code = http.StatusBadRequest
		m.Message = "no hay conincidencia con el identificador"
		return commons.DisplayMessage(c, &m)
	}
	u.ID = id
	controllers.UserUpdate(u, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserDeleteRoutes Creacion de usuario
func SetUserDeleteRoutes(c echo.Context) error {
	m := models.Message{}
	getUserInterface(c, &m.User)
	var u models.User
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	u.ID = id
	controllers.UserDelete(u, &m)
	return commons.DisplayMessage(c, &m)
}
