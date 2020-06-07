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

//SetGetUserRoutes Creacion de usuario
func SetGetUserRoutes(c echo.Context) error {
	user := c.Get(commons.User)
	m := models.Message{}

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

	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, &m.User)
	controllers.UserGet(u, &m)
	return commons.DisplayMessage(c, &m)
}

//SetEditUserRoutes Creacion de usuario
func SetEditUserRoutes(c echo.Context) error {
	user := c.Get(commons.User)
	m := models.Message{}

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
	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, &m.User)
	controllers.UserUpdate(u, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDeleteUserRoutes Creacion de usuario
func SetDeleteUserRoutes(c echo.Context) error {
	user := c.Get(commons.User)
	m := models.Message{}
	var u models.User
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	u.ID = id
	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, &m.User)
	controllers.UserDelete(u, &m)
	return commons.DisplayMessage(c, &m)
}
