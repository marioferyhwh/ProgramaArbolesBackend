package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

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
	flag.IntVar(&commons.Port, "port", 0, " puerto del servidor")
	var err error
	s := os.Getenv("PORT")
	fmt.Println("port:", s)
	if commons.Port == 0 {
		port, err := strconv.ParseInt(s, 10, 64)

		fmt.Println("port2:", port)
		if err != nil {
			commons.Port = int(port)
		} else {
			commons.Port = 8080
		}
	}

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
		fmt.Printf("error leyendo las direcciones ip del hostname: %v\n", err)
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
		AllowHeaders: []string{"*"},
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
