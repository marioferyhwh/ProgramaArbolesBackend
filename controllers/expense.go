package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createExpense crea gasto con una conexion ya existente
func createExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Create(e).Error
	return err
}

//getExpense trae gasto con una conexion ya existente
func getExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(e).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteExpense se borra el gasto con una conexion ya existente
func deleteExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(e).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateExpense se borra el gasto con una conexion ya existente
func updateExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(e).Omit(omitList...).Updates(e).Error
	return err
}

/*······························································
································································
··············· listado de descripcion de gasto por cada collection
································································
······························································*/

//createExpenseDescrip crea descripcion de gasto con una conexion ya existente
func createExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Create(ed).Error
	return err
}

//getExpenseDescrip trae descripcion de gasto con una conexion ya existente
func getExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(ed).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteExpenseDescrip se borra el descripcion de gasto con una conexion ya existente
func deleteExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ed).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateExpenseDescrip se borra el descripcion de gasto con una conexion ya existente
func updateExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ed).Omit(omitList...).Updates(ed).Error
	return err
}
