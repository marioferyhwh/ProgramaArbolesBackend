package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createLoan crea Prestamos con una conexion ya existente
func createLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Create(l).Error
	return err
}

//getLoan trae Prestamos con una conexion ya existente
func getLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(l).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteLoan se borra el Prestamos con una conexion ya existente
func deleteLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(l).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateLoan se borra el Prestamos con una conexion ya existente
func updateLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(l).Omit(omitList...).Updates(l).Error
	return err
}

/*······························································
································································
··············· estado de prestamos
································································
······························································*/

//createLoanState crea estado de  prestamos con una conexion ya existente
func createLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Create(ls).Error
	return err
}

//getLoanState trae estado de  prestamos con una conexion ya existente
func getLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(ls).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteLoanState se borra el estado de  prestamos con una conexion ya existente
func deleteLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ls).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateLoanState se borra el estado de  prestamos con una conexion ya existente
func updateLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ls).Omit(omitList...).Updates(ls).Error
	return err
}

/*······························································
································································
··············· pagos de los prestamos
································································
······························································*/

//createLoanPayment crea pagos de prestamos con una conexion ya existente
func createLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	err := db.Create(lp).Error
	return err
}

//getLoanPayment trae pagos de prestamos con una conexion ya existente
func getLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(lp).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteLoanPayment se borra el pagos de prestamos con una conexion ya existente
func deleteLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(lp).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateLoanPayment se borra el pagos de prestamos con una conexion ya existente
func updateLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(lp).Omit(omitList...).Updates(lp).Error
	return err
}
