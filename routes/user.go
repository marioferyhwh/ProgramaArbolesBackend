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

//getParams64
func getParams64(c echo.Context) (uint64, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(i64), nil
}

//getParams32
func getParams32(c echo.Context) (uint32, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i64), nil
}

//getParams16
func getParams16(c echo.Context) (uint16, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint16(i64), nil
}

//getParams8
func getParams8(c echo.Context) (uint8, error) {
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint8(i64), nil
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

//SetUserGetRoutes trae de usuario
func SetUserGetRoutes(c echo.Context) error {
	m := models.Message{}
	getUserInterface(c, &m.User)
	var u models.User
	id, err := getParams32(c)
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

//SetUserEditRoutes editar de usuario
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
	id, err := getParams32(c)
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

//SetUserDeleteRoutes borarr usuario
func SetUserDeleteRoutes(c echo.Context) error {
	m := models.Message{}
	getUserInterface(c, &m.User)
	var u models.User
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	u.ID = id
	controllers.UserDelete(u, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· telefono de usuario
································································
······························································*/

//SetUserTelCreateRoutes Creacion de telefono de usuario
func SetUserTelCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ut models.UserTel
	err := c.Bind(&ut)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego telefono de usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.UserTelCreate(ut, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserTelGetRoutes traer telefono de usuario
func SetUserTelGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ut models.UserTel
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	ut.ID = id
	controllers.UserTelGet(ut, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserTelGetListRoutes listado de ubicaiones
func SetUserTelGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ut models.UserTel
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	ut.CodUser = uint16(id)
	controllers.UserTelGetList(ut, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserTelEditRoutes actualizar telefono de usuario
func SetUserTelEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ut models.UserTel
	err := c.Bind(&ut)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego telefono de usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	ut.ID, err = getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	controllers.UserTelUpdate(ut, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserTelDeleteRoutes borrar telefono de usuario
func SetUserTelDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ut models.UserTel
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de telefono de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	ut.ID = id
	controllers.UserTelDelete(ut, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· nivel de usuario
································································
······························································*/

//SetUserLevelCreateRoutes Creacion de nivel de usuario
func SetUserLevelCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserLevel
	err := c.Bind(&uc)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego nivel de usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.UserLevelCreate(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserLevelGetRoutes traer nivel de usuario
func SetUserLevelGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserLevel
	id, err := getParams8(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de nivel de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserLevelGet(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserLevelGetListRoutes listado de ubicaiones
func SetUserLevelGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserLevel
	controllers.UserLevelGetList(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserLevelEditRoutes actualizar nivel de usuario
func SetUserLevelEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserLevel
	err := c.Bind(&uc)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego nivel de usuario ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams8(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de nivel de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserLevelUpdate(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserLevelDeleteRoutes borrar nivel de usuario
func SetUserLevelDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserLevel
	id, err := getParams8(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de nivel de usuario no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserLevelDelete(uc, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· usuario en cobro
································································
······························································*/

//SetUserCollectionCreateRoutes Creacion de asignacion de usuario a cobro
func SetUserCollectionCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserCollection
	err := c.Bind(&uc)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego asignacion de usuario a cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.UserCollectionCreate(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserCollectionGetRoutes traer asignacion de usuario a cobro
func SetUserCollectionGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserCollection
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de asignacion de usuario a cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserCollectionGet(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserCollectionGetListRoutes listado de ubicaiones
func SetUserCollectionGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserCollection
	id, err := getParams32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de asignacion de usuario a cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.CodCollection = id
	controllers.UserCollectionGetList(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserCollectionEditRoutes actualizar asignacion de usuario a cobro
func SetUserCollectionEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserCollection
	err := c.Bind(&uc)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego asignacion de usuario a cobro ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de asignacion de usuario a cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserCollectionUpdate(uc, &m)
	return commons.DisplayMessage(c, &m)
}

//SetUserCollectionDeleteRoutes borrar asignacion de usuario a cobro
func SetUserCollectionDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var uc models.UserCollection
	id, err := getParams64(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de asignacion de usuario a cobro no valido"
		return commons.DisplayMessage(c, &m)
	}
	uc.ID = id
	controllers.UserCollectionDelete(uc, &m)
	return commons.DisplayMessage(c, &m)
}
