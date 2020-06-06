package controllers

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createLoan crea Prestamos con una conexion ya existente
func createLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Create(l).Error
	return err
}

//getLoan trae Prestamos con una conexion ya existente
func getLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client,cod_collection,cod_user").First(l).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanList trae Prestamos con una conexion ya existente
func getLoanList(l *[]models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client").Find(l).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoan se borra el Prestamos con una conexion ya existente
func updateLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "initial_value", "interest", "quota", "cod_client", "cod_collection", "cod_user"}
	err := db.Model(l).Omit(omitList...).Updates(l).Error
	return err
}

//deleteLoan se borra el Prestamos con una conexion ya existente
func deleteLoan(l *models.Loan, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(l).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· estado de prestamos
································································
······························································*/

//LoanStateCreate crea un nuevo tipo de documento
func LoanStateCreate(ls models.LoanState, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createLoanState(&ls, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de prestamo no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de prestamo creado"
	m.Data = ls
}

//createLoanState crea estado de  prestamos con una conexion ya existente
func createLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Create(ls).Error
	return err
}

//getLoanState trae estado de  prestamos con una conexion ya existente
func getLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,state").First(ls).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanStateList trae estado de  prestamos con una conexion ya existente
func getLoanStateList(ls *[]models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,state").Find(ls).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoanState se borra el estado de  prestamos con una conexion ya existente
func updateLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ls).Omit(omitList...).Updates(ls).Error
	return err
}

//deleteLoanState se borra el estado de  prestamos con una conexion ya existente
func deleteLoanState(ls *models.LoanState, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ls).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
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
	err := db.Select("id,created_at,updated_at,cod_loan,cash,cod_user").First(lp).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanPaymentList trae pagos de prestamos con una conexion ya existente
func getLoanPaymentList(lp *[]models.LoanPayment, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,updated_at,cod_loan,cash,cod_user").Find(lp).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoanPayment se borra el pagos de prestamos con una conexion ya existente
func updateLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection", "cod_user", "cod_loan"}
	err := db.Model(lp).Omit(omitList...).Updates(lp).Error
	return err
}

//deleteLoanPayment se borra el pagos de prestamos con una conexion ya existente
func deleteLoanPayment(lp *models.LoanPayment, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(lp).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}
