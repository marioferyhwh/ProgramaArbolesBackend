package routes

import "github.com/gorilla/mux"

//SetLoginRoutes inicio de seccion
func SetLoginRoutes(router *mux.Router) {
	router.HandleFunc("/api/login", con)
}
