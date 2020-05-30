package controllers

import (
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//Login funcion de inicio de seccion
func Login(w http.ResponseWriter, r *http.Request) {

	m := models.Message{
		Code:    200,
		Message: "login",
	}

	commons.DisplayMessage(w, m)
}
