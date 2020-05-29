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
	db.DropTableIfExists(&models.UserLevel{})
	db.DropTableIfExists(&models.LoanState{})
	db.DropTableIfExists(&models.DocumentType{})

	db.CreateTable(&models.Collection{})
	db.CreateTable(&models.UserLevel{})
	db.CreateTable(&models.LoanState{})
	db.CreateTable(&models.DocumentType{})
}
