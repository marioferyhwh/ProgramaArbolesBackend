package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

/*······························································
································································
··············· gasto
································································
······························································*/

//ExpenseCreate crea un nuevo gasto
func ExpenseCreate(e models.Expense, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createExpense(&e, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "gasto no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "gasto creado"
	m.Data = e
}

//ExpenseGet traer un nuevo gasto
func ExpenseGet(e models.Expense, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getExpense(&e, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro gasto"
		return
	}
	m.Code = http.StatusOK
	m.Message = "gasto creado"
	m.Data = e
}

//ExpenseGetList traer lista de gasto
func ExpenseGetList(e models.Expense, m *models.Message) {
	es := []models.Expense{e}
	db := configuration.GetConnection()
	defer db.Close()
	err := getExpenseList(&es, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de gastos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de gastos"
	m.Data = es
}

//ExpenseUpdate se edita un gasto
func ExpenseUpdate(e models.Expense, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateExpense(&e, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "gasto no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo gasto"
	m.Data = e
}

//ExpenseDelete se borra un gasto
func ExpenseDelete(e models.Expense, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteExpense(&e, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "gasto no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = e
}

/*······························································
······························································*/

//createExpense crea gasto
func createExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Create(e).Error
	return err
}

//getExpense trae gasto
func getExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,cash,cod_expense_descrip,cod_user,cod_collection").First(e).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getExpenseList trae gasto
func getExpenseList(es *[]models.Expense, m *models.Message, db *gorm.DB) error {
	var e models.Expense
	if len(*es) == 0 {
		e = (*es)[0]
	}
	where := fmt.Sprintf("cod_collection = %v", e.CodCollection)
	if e.CodUser != 0 {
		where += fmt.Sprintf("cod_user = %v", e.CodUser)
	}
	err := db.Where(where).Select("id,cash,cod_expense_descrip,cod_user").Find(es).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateExpense actualizar el gasto
func updateExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "cod_expense_descrip", "cod_user", "cod_collection"}
	err := db.Model(e).Omit(omitList...).Updates(e).Error
	return err
}

//deleteExpense se borra el gasto
func deleteExpense(e *models.Expense, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(e).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· listado de descripcion de gasto por cada collection
································································
······························································*/

//ExpenseDescripCreate crea un nuevo descripcion de gastos
func ExpenseDescripCreate(ed models.ExpenseDescrip, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createExpenseDescrip(&ed, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de gastos no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion de gastos creado"
	m.Data = ed
}

//ExpenseDescripGet traer un nuevo descripcion de gastos
func ExpenseDescripGet(ed models.ExpenseDescrip, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getExpenseDescrip(&ed, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro descripcion de gastos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion de gastos creado"
	m.Data = ed
}

//ExpenseDescripGetList traer lista de descripcion de gastos
func ExpenseDescripGetList(ed models.ExpenseDescrip, m *models.Message) {
	eds := []models.ExpenseDescrip{ed}
	db := configuration.GetConnection()
	defer db.Close()
	err := getExpenseDescripList(&eds, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de descripcion de gastoss"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de descripcion de gastoss"
	m.Data = eds
}

//ExpenseDescripUpdate se edita un descripcion de gastos
func ExpenseDescripUpdate(ed models.ExpenseDescrip, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateExpenseDescrip(&ed, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de gastos no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo descripcion de gastos"
	m.Data = ed
}

//ExpenseDescripDelete se borra un descripcion de gastos
func ExpenseDescripDelete(ed models.ExpenseDescrip, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteExpenseDescrip(&ed, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de gastos no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = ed
}

/*······························································
······························································*/

//createExpenseDescrip crea descripcion de gasto
func createExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Create(ed).Error
	return err
}

//getExpenseDescrip trae descripcion de gasto
func getExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,cod_collection,descrip").First(ed).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getExpenseDescripList trae descripcion de gasto
func getExpenseDescripList(eds *[]models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	var ed models.ExpenseDescrip
	if len(*eds) == 1 {
		ed = (*eds)[0]
	}
	err := db.Where("cod_collection = ?", ed.CodCollection).Select("id,descrip").Find(eds).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateExpenseDescrip se borra el descripcion de gasto
func updateExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection"}
	err := db.Model(ed).Omit(omitList...).Updates(ed).Error
	return err
}

//deleteExpenseDescrip se borra el descripcion de gasto
func deleteExpenseDescrip(ed *models.ExpenseDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ed).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
