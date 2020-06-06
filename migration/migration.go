package migration

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//Migrate Crea Tablas
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()
	// deleteTables(db)
	// createTables(db)
	// createConstrain(db)
	// createDataInit()

	var m models.Message
	controllers.BusinessTypeGetList(models.BusinessType{}, &m)

	switch m.Data.(type) {

	case []models.BusinessType:
		b := m.Data.([]models.BusinessType)
		for _, d := range b {
			fmt.Println(d)
		}

	default:
		fmt.Println(m.Data)

	}
	// var user models.User
	// db.Find(&user)
	// var documents []models.DocumentType
	// db.Debug().Model(&user).Related(&documents, "DocumentType")
	// var user1 models.User
	// db.Debug().Where("id = 2").Find(&user1)
	// db.Model(&user1).Related(&documents, "DocumentType")
	// user1.DocumentType = documents[0]

}

func createDataInit() {
	m := models.Message{}
	documents := []models.DocumentType{
		{
			Descrip:   "Targeta de iendtificacion",
			NameShort: "TI",
		},
		{
			Descrip:   "Cedula De Ciudadania",
			NameShort: "CC",
		},
		{
			Descrip:   "DNI",
			NameShort: "DNI",
		},
	}
	users := []models.User{
		{
			NickName:        "forest",
			Email:           "forest",
			Password:        "forest",
			CodDocumentType: "CC",
			Document:        "1111111111",
			Name:            "forest",
			Admin:           true,
		},
	}
	m.User = users[0]
	businessTypes := []models.BusinessType{
		{Descrip: "mercado"},
		{Descrip: "joyeria"},
		{Descrip: "banco"},
	}
	telDescrips := []models.TelDescrip{
		{Descrip: "personal"},
		{Descrip: "trabajo"},
		{Descrip: "casa"},
		{Descrip: "otro"},
	}
	loanStates := []models.LoanState{
		{State: "personal"},
		{State: "trabajo"},
		{State: "casa"},
		{State: "otro"},
	}

	for _, d := range documents {
		controllers.DocumentTypeCreate(d, &m)
	}

	for _, bt := range businessTypes {
		controllers.BusinessTypeCreate(bt, &m)
	}

	for _, td := range telDescrips {
		controllers.TelDescripCreate(td, &m)
	}

	for _, ls := range loanStates {
		controllers.LoanStateCreate(ls, &m)
	}
	for _, u := range users {
		controllers.UserCreate(u, &m)
	}
	fmt.Println(m)
}

func deleteTables(db *gorm.DB) {
	db.DropTableIfExists(&models.Expense{})
	db.DropTableIfExists(&models.ExpenseDescrip{})
	db.DropTableIfExists(&models.CollectionCash{})
	db.DropTableIfExists(&models.LoanPayment{})
	db.DropTableIfExists(&models.Loan{})
	db.DropTableIfExists(&models.ClientTel{})
	db.DropTableIfExists(&models.Client{})
	db.DropTableIfExists(&models.UserTel{})
	db.DropTableIfExists(&models.UserCollection{})
	db.DropTableIfExists(&models.ClientListLocation{})
	db.DropTableIfExists(&models.BusinessType{})
	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.TelDescrip{})
	db.DropTableIfExists(&models.DocumentType{})
	db.DropTableIfExists(&models.LoanState{})
	db.DropTableIfExists(&models.UserLevel{})
	db.DropTableIfExists(&models.Collection{})
}

func createTables(db *gorm.DB) {
	db.CreateTable(&models.Collection{})
	db.CreateTable(&models.UserLevel{})
	db.CreateTable(&models.LoanState{})
	db.CreateTable(&models.DocumentType{})
	db.CreateTable(&models.TelDescrip{})
	db.CreateTable(&models.BusinessType{})
	db.CreateTable(&models.ClientListLocation{})
	db.CreateTable(&models.User{})
	db.CreateTable(&models.UserCollection{})
	db.CreateTable(&models.UserTel{})
	db.CreateTable(&models.Client{})
	db.CreateTable(&models.ClientTel{})
	db.CreateTable(&models.Loan{})
	db.CreateTable(&models.LoanPayment{})
	db.CreateTable(&models.CollectionCash{})
	db.CreateTable(&models.ExpenseDescrip{})
	db.CreateTable(&models.Expense{})
}

func createConstrain(db *gorm.DB) {
	//db.Model(&models.DocumentType{}).AddUniqueIndex("document_types_name_short_key", "name_short")

	//db.Model(&models.User{}).AddUniqueIndex("users_cdocumentt_document_key", "cod_document_type", "document").AddForeignKey("cod_document_type", "document_types(name_short)", "restrict", "restrict")

	// db.Model(&models.ListLocation{}).AddForeignKey("cod_collection", "collections(id)", "restrict", "restrict").AddUniqueIndex("list_locations_ccollection_descrip_key", "cod_collection", "descrip")
	// db.Model(&models.ListUser{}).AddForeignKey("cod_user", "users(id)", "restrict", "restrict").AddForeignKey("cod_collection", "collections(id)", "restrict", "restrict").AddForeignKey("cod_user_level", "user_levels(id)", "restrict", "restrict").AddUniqueIndex("list_users_cuser_ccollection_key", "cod_user", "cod_collection").AddIndex("list_users_ccollection_cuser_key", "cod_collection", "cod_user")
	// db.Model(&models.UserTel{}).AddForeignKey("cod_user", "users(id)", "restrict", "restrict").AddForeignKey("cod_tel_descrip", "tel_descrips(id)", "restrict", "restrict")
	// db.Model(&models.Client{}).AddUniqueIndex("clients_cdocumentt_document_key", "cod_document_type", "document").AddForeignKey("cod_document_type", "document_types(id)", "restrict", "restrict").AddForeignKey("cod_collection", "collections(id)", "restrict", "restrict").AddForeignKey("cod_loan_state", "loan_states(id)", "restrict", "restrict").AddForeignKey("cod_business_type", "business_types(id)", "restrict", "restrict").AddForeignKey("cod_list_location", "list_locations(id)", "restrict", "restrict").AddForeignKey("cod_user", "users(id)", "restrict", "restrict")
	fmt.Println("se crea contraint")

}
