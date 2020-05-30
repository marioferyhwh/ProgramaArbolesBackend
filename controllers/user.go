package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

//Login funcion de inicio de seccion
func Login(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal("login")
	if err != nil {
		log.Fatal("no se combirtio mensaje->", err)
	}
	w.Write(j)
}
