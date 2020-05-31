package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

type model struct {
	Email    string
	Password string
}

//Login funcion de inicio de seccion
func Login(c echo.Context) error {
	var user models.User
	m := models.Message{}

	err := c.Bind(&user)

	if err != nil {
		fmt.Printf("Error : %s\n", err)
		m.Code = http.StatusNotFound
		m.Message = fmt.Sprint("no", err.Error())
		return commons.DisplayMessage(c, &m)
	}

	pwd := encriptPasswordUser(user.Password)

	db := configuration.GetConnection()
	defer db.Close()
	db.Where("(nick_name = ? or email = ?) and password = ?", user.Email, user.Email, pwd).First(&user)
	user.Password = ""

	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y clave"
		return commons.DisplayMessage(c, &m)
	}
	// user.Password = ""
	token := commons.GenetateJWT(user)
	/*
		fmt.Print(token)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = err.Error()
		return
		}
	*/
	m.Code = http.StatusOK
	m.NewToken = token
	return commons.DisplayMessage(c, &m)
}

func encriptPasswordUser(password string) string {
	c := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", c)
}

//UserCreateInter crea el usuario
func UserCreateInter(user models.User, m *models.Message) {
	pwd := encriptPasswordUser(user.Password)
	user.Password = pwd
	db := configuration.GetConnection()
	defer db.Close()

	err := db.Create(&user).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("error a crear el registro :%s", err)
		return
	}
	m.Code = http.StatusOK
	m.Message = "Usuatio Creado"
}

//UserCreate Creacion de usuario
func UserCreate(c echo.Context) error {
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
	UserCreateInter(user, &m)
	return commons.DisplayMessage(c, &m)

}
