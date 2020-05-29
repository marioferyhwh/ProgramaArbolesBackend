package migration

import (
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
)

//Migrate Crea Tablas
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	/*
				db.DropTableIfExists(&models.User{})

		   	db.DropTableIfExists(&models.Collection{})
		   	db.DropTableIfExists(&models.UserLevel{})
		   	db.DropTableIfExists(&models.LoanState{})
		   	db.DropTableIfExists(&models.DocumentType{})

		   	db.CreateTable(&models.Collection{})
		   	db.CreateTable(&models.UserLevel{})
		   	db.CreateTable(&models.LoanState{})
		   	db.CreateTable(&models.DocumentType{})
		   	db.CreateTable(&models.User{})
				 db.Model(&models.User{}).AddForeignKey("cod_document_type", "document_types(id)", "restrict", "restrict")
	*/
}
