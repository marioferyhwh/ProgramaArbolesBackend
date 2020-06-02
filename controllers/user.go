package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

type model struct {
	Email    string
	Password string
}

//encriptPasswordUser Se encripta la clave
func encriptPasswordUser(password string) string {
	c := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", c)
}

//Login funcion de inicio de seccion
func Login(user models.User, m *models.Message) {
	pwd := encriptPasswordUser(user.Password)

	db := configuration.GetConnection()
	defer db.Close()
	db.Where("(nick_name = ? or email = ?) and password = ?", user.Email, user.Email, pwd).First(&user)
	user.Password = ""
	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y/o clave"
		return
	}
	// user.Password = ""
	token, err := commons.GenetateJWT(user)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "error generando token"
		return
	}
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
}

//UserCreate crea el usuario
func UserCreate(user models.User, m *models.Message) {
	//validacion de datos de usuario
	if user.Email == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta email"
		return
	}
	if user.Active {
		user.Active = true
	}
	if user.NickName == "" {
		user.NickName = user.Email
	}
	if user.CodDocumentType == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta tipo de documento"
		return
	}
	if user.Document == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta Documento"
		return
	}
	if user.Password == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta clave"
		return
	}
	//encriptar clave
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

	user.Password = ""
	user.ConfirmPassword = ""

	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}

//GetUser trae un usario
func GetUser(user models.User, m *models.Message) {

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("id = ?", user.ID).First(&user)
	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y/o clave"
		return
	}
	user.Password = ""
	user.ConfirmPassword = ""

	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}

//EditUser trae un usario
func EditUser(user models.User, m *models.Message) {

	user.Password = ""
	user.ConfirmPassword = ""

	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}

//DeleteUser trae un usario
func DeleteUser(user models.User, m *models.Message) {

	user.Password = ""
	user.ConfirmPassword = ""

	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}
