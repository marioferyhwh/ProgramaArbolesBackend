package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

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

	hostName, err := os.Hostname()
	if err != nil {
		log.Println("error leyendo el host name")
		return
	}
	addrs, err := net.LookupHost(hostName)
	if err != nil {
		fmt.Printf("erroro leyendo las direcciones ip del hostname: %v\n", err)
		return
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
	for _, a := range addrs {
		log.Println("direccion:", a)
	}
	log.Println("puerto:", commons.Port)
	log.Panicln(server.ListenAndServe())
	log.Println("servicion detenido ¿eso es lo que queria?")
	// fmt.Println("Finaliza Backend")
}
