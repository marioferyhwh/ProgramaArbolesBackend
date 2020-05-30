package migration

import (
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//Migrate Crea Tablas
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.DropTableIfExists(&models.UserTel{})

	db.DropTableIfExists(&models.ListUser{})
	db.DropTableIfExists(&models.ListLocation{})

	db.DropTableIfExists(&models.BusinessTypes{})
	db.DropTableIfExists(&models.TelDescrip{})
	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Collection{})
	db.DropTableIfExists(&models.UserLevel{})
	db.DropTableIfExists(&models.LoanState{})
	db.DropTableIfExists(&models.DocumentType{})
	db.CreateTable(&models.Collection{})
	db.CreateTable(&models.UserLevel{})
	db.CreateTable(&models.LoanState{})
	db.CreateTable(&models.DocumentType{})
	db.CreateTable(&models.TelDescrip{})
	db.CreateTable(&models.User{})

	db.Model(&models.User{}).AddForeignKey("cod_document_type", "document_types(id)", "restrict", "restrict")

	db.CreateTable(&models.BusinessTypes{})

	db.CreateTable(&models.ListLocation{})
	db.Model(&models.ListLocation{}).AddForeignKey("cod_collection", "collections(id)", "restrict", "restrict").AddUniqueIndex("list_locations_ccollection_descrip_key", "cod_collection", "descrip")

	db.CreateTable(&models.ListUser{})
	db.Model(&models.ListUser{}).AddForeignKey("cod_user", "users(id)", "restrict", "restrict").AddForeignKey("cod_collection", "collections(id)", "restrict", "restrict").AddForeignKey("cod_user_level", "user_levels(id)", "restrict", "restrict").AddUniqueIndex("list_users_uk_cuser_ccollection", "cod_user", "cod_collection").AddIndex("list_users_ccollection_cuser_key", "cod_collection", "cod_user")

	db.CreateTable(&models.UserTel{})
	db.Model(&models.UserTel{}).AddForeignKey("cod_user", "users(id)", "restrict", "restrict").AddForeignKey("cod_tel_descrip", "tel_descrips(id)", "restrict", "restrict")
}
