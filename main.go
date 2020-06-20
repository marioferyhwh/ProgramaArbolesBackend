package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/migration"
	"github.com/marioferyhwh/IMFBackend_forest/routes"
)

func main() {
	log.Println("Inicia Backend")
	testdb := false
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

	if testdb {
		fmt.Println("se termina ejecucuion")
		return
	}

	//middleware
	//routes
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	routes.InitRoutes(e)
	//inicio del servidor
	for _, a := range addrs {
		log.Println("direccion:", a)
	}
	log.Println("puerto:", commons.Port)
	err = e.Start(fmt.Sprintf(":%d", commons.Port))
	if err != nil {
		log.Println("servicion detenido ¿eso es lo que queria?")
		log.Println(err)
	}
}
