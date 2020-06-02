package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetLoginRoutes inicio de seccion
func SetLoginRoutes(c echo.Context) error {
	var user models.User
	m := models.Message{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
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
	//defer
	err := c.Bind(&user)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usario ->", err)
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
	user := models.User{}
	m := models.Message{}
	//defer
	err := c.Bind(&user)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	if user.ConfirmPassword != user.Password {
		m.Code = http.StatusBadRequest
		m.Message = "contraseñas no coninciden"
		return commons.DisplayMessage(c, &m)
	}
	controllers.GetUser(user, &m)
	return commons.DisplayMessage(c, &m)
}

//SetEditUserRoutes Creacion de usuario
func SetEditUserRoutes(c echo.Context) error {
	user := models.User{}
	m := models.Message{}
	//defer
	err := c.Bind(&user)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	if user.ConfirmPassword != user.Password {
		m.Code = http.StatusBadRequest
		m.Message = "contraseñas no coninciden"
		return commons.DisplayMessage(c, &m)
	}
	controllers.EditUser(user, &m)
	return commons.DisplayMessage(c, &m)
}

//SetDeleteUserRoutes Creacion de usuario
func SetDeleteUserRoutes(c echo.Context) error {
	user := models.User{}
	m := models.Message{}
	//defer
	err := c.Bind(&user)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego usario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	if user.ConfirmPassword != user.Password {
		m.Code = http.StatusBadRequest
		m.Message = "contraseñas no coninciden"
		return commons.DisplayMessage(c, &m)
	}
	controllers.DeleteUser(user, &m)
	return commons.DisplayMessage(c, &m)
}
