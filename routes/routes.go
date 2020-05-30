package routes

import "github.com/gorilla/mux"

//InitRoutes en rutamiento de la API
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	return router

}
