package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/migration"
	"github.com/marioferyhwh/IMFBackend_forest/routes"
	"github.com/urfave/negroni"
)

func main() {
	log.Println("Inicia Backend")
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generar la migracion")
	flag.IntVar(&commons.Port, "port", 8080, " puerto del servidor")
	flag.Parse()
	if migrate == "yes" {
		migration.Migrate()
	}

	router := routes.InitRoutes()
	//middleware
	n := negroni.Classic()
	n.UseHandler(router)

	//inicio del servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}
	log.Println("Inicia servidor en el puerto:", commons.Port)
	log.Panicln(server.ListenAndServe())
	log.Println("servicion detenido ¿eso es lo que queria?")
	// fmt.Println("Finaliza Backend")
}
