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

//Login funcion de inicio de seccion
func Login(user models.User, m *models.Message) {
	pwd := encriptPasswordUser(user.Password)

	db := configuration.GetConnection()
	defer db.Close()
	db.Where("(nick_name = ? or email = ?) and password = ? and active = 1", user.Email, user.Email, pwd).First(&user)
	user.Password = ""
	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y/o clave"
		return
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
}

func encriptPasswordUser(password string) string {
	c := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", c)
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
