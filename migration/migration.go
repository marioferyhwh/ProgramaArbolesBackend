package migration

import (
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//Migrate Crea Tablas
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()
	db.DropTableIfExists(&models.Collection{})

	db.CreateTable(&models.Collection{})
}
