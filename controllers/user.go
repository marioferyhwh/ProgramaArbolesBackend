package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//Login funcion de inicio de seccion
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}
	defer commons.DisplayMessage(w, m)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Code = http.StatusNotFound
		m.Message = err.Error()
		return
	}

	pwd := encriptPasswordUser(user.Password)

	db := configuration.GetConnection()
	db.Where("(nick_name = ? or email = ?) and password = ?", user.Email, user.Email, pwd).First(&user)

	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y clave"
		return
	}

	m.Code = http.StatusOK
	m.Message = ""
}

func encriptPasswordUser(password string) string {
	c := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", c)
}
