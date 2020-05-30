package routes

import (
	"github.com/gorilla/mux"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
)

//SetLoginRoutes inicio de seccion
func SetLoginRoutes(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
