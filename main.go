package main

import (
	"flag"
	"fmt"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/migration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

func main() {
	fmt.Println("Inicia Backend")
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generar la migracion")
	flag.IntVar(&commons.Port, "port", 8080, " puerto del servidor")
	flag.Parse()
	if migrate == "yes" {
		migration.Migrate()
	}
	var user models.User

	fmt.Println(user)
	fmt.Println("Finaliza Backend")
}
