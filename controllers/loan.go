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
··············· prestamo
································································
······························································*/

//LoanCreate crea un nuevo prestamo
func LoanCreate(l models.Loan, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createLoan(&l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "prestamo no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "prestamo creado"
	m.Data = l
}

//LoanGet traer un nuevo prestamo
func LoanGet(l models.Loan, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoan(&l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro prestamo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "prestamo creado"
	m.Data = l
}

//LoanGetList traer lista de prestamo
func LoanGetList(l models.Loan, m *models.Message) {
	ls := []models.Loan{l}
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanList(&ls, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de prestamos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de prestamos"
	m.Data = ls
}

//LoanUpdate se edita un prestamo
func LoanUpdate(l models.Loan, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateLoan(&l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "prestamo no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo prestamo"
	m.Data = l
}

//LoanDelete se borra un prestamo
func LoanDelete(l models.Loan, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteLoan(&l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "prestamo no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = l
}

/*······························································
······························································*/

//createLoan crea Prestamos
func createLoan(l *models.Loan, db *gorm.DB) error {
	err := db.Create(l).Error
	return err
}

//getLoan trae Prestamos
func getLoan(l *models.Loan, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client,cod_collection,cod_user").First(l).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanList trae Prestamos
func getLoanList(ls *[]models.Loan, db *gorm.DB) error {
	var l models.Loan
	if len(*ls) == 1 {
		l = (*ls)[0]
	}
	where := ""
	if l.CodCollection != 0 {
		where = fmt.Sprintf("cod_collextion = %v", l.CodCollection)
	}
	if l.CodClient != 0 {
		where = fmt.Sprintf("cod_client = %v", l.CodClient)
	} else if l.CodLoanState != 0 {
		where += fmt.Sprintf(" and cod_loan_state = %v", l.CodLoanState)
	}
	err := db.Where(where).Select("id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client").Find(ls).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoan se borra el Prestamos
func updateLoan(l *models.Loan, db *gorm.DB) error {
	omitList := []string{"id", "initial_value", "interest", "quota", "cod_client", "cod_collection", "cod_user"}
	err := db.Model(l).Omit(omitList...).Updates(l).Error
	return err
}

//deleteLoan se borra el Prestamos
func deleteLoan(l *models.Loan, db *gorm.DB) error {
	err := db.Unscoped().Delete(l).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· estado de prestamos
································································
······························································*/

//LoanStateCreate crea un nuevo estado de prestamo
func LoanStateCreate(ls models.LoanState, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createLoanState(&ls, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de prestamo no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de prestamo creado"
	m.Data = ls
}

//LoanStateGet traer un nuevo estado de prestamo
func LoanStateGet(ls models.LoanState, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanState(&ls, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro estado de prestamo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de prestamo creado"
	m.Data = ls
}

//LoanStateGetList traer lista de estado de prestamo
func LoanStateGetList(ls models.LoanState, m *models.Message) {
	lss := []models.LoanState{ls}
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanStateList(&lss, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de estado de prestamos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de estado de prestamos"
	m.Data = lss
}

//LoanStateUpdate se edita un estado de prestamo
func LoanStateUpdate(ls models.LoanState, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateLoanState(&ls, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de prestamo no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo estado de prestamo"
	m.Data = ls
}

//LoanStateDelete crea un nuevo estado de prestamo
func LoanStateDelete(lp models.LoanState, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteLoanState(&lp, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "estado de prestamo no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "estado de prestamo se borro"
	m.Data = lp
}

/*······························································
······························································*/

//createLoanState crea estado de  prestamos
func createLoanState(ls *models.LoanState, db *gorm.DB) error {
	err := db.Create(ls).Error
	return err
}

//getLoanState trae estado de  prestamos
func getLoanState(ls *models.LoanState, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,state").First(ls).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanStateList trae estado de  prestamos
func getLoanStateList(lss *[]models.LoanState, db *gorm.DB) error {
	err := db.Select("id,state").Find(lss).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoanState se borra el estado de  prestamos
func updateLoanState(ls *models.LoanState, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ls).Omit(omitList...).Updates(ls).Error
	return err
}

//deleteLoanState se borra el estado de  prestamos
func deleteLoanState(ls *models.LoanState, db *gorm.DB) error {
	err := db.Unscoped().Delete(ls).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· pagos de los prestamos
································································
······························································*/

//LoanPaymentCreate crea un nuevo pago a prestamo
func LoanPaymentCreate(lp models.LoanPayment, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createLoanPayment(&lp, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "pago a prestamo no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "pago a prestamo creado"
	m.Data = lp
}

//LoanPaymentGet traer un nuevo pago a prestamo
func LoanPaymentGet(lp models.LoanPayment, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanPayment(&lp, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro pago a prestamo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "pago a prestamo creado"
	m.Data = lp
}

//LoanPaymentGetList traer lista de pago a prestamo
func LoanPaymentGetList(lp models.LoanPayment, m *models.Message) {
	lps := []models.LoanPayment{lp}
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanPaymentList(&lps, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de pago a prestamos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de pago a prestamos"
	m.Data = lps
}

//LoanPaymentUpdate se edita un pago a prestamo
func LoanPaymentUpdate(lp models.LoanPayment, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateLoanPayment(&lp, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "pago a prestamo no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo pago a prestamo"
	m.Data = lp
}

//LoanPaymentDelete se borra un pago a prestamo
func LoanPaymentDelete(lp models.LoanPayment, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteLoanPayment(&lp, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "pago a prestamo no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = lp
}

/*······························································
······························································*/

//createLoanPayment crea pagos de prestamos
func createLoanPayment(lp *models.LoanPayment, db *gorm.DB) error {
	err := db.Create(lp).Error
	return err
}

//getLoanPayment trae pagos de prestamos
func getLoanPayment(lp *models.LoanPayment, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,cod_loan,cash,cod_user").First(lp).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanPaymentList trae pagos de prestamos
func getLoanPaymentList(lps *[]models.LoanPayment, db *gorm.DB) error {
	var lp models.LoanPayment
	if len(*lps) == 1 {
		lp = (*lps)[0]
	}
	where := fmt.Sprintf("cod_collection = %v", lp.CodCollection)
	if lp.CodLoan != 0 {
		where += fmt.Sprintf(" and cod_loan = %v", lp.CodLoan)
	} else if lp.CodUser != 0 {
		where += fmt.Sprintf(" and cod_user = %v", lp.CodUser)
	}
	err := db.Where(where).Select("id,updated_at,cod_loan,cash,cod_user").Find(lps).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoanPayment se borra el pagos de prestamos
func updateLoanPayment(lp *models.LoanPayment, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection", "cod_user", "cod_loan"}
	err := db.Model(lp).Omit(omitList...).Updates(lp).Error
	return err
}

//deleteLoanPayment se borra el pagos de prestamos
func deleteLoanPayment(lp *models.LoanPayment, db *gorm.DB) error {
	err := db.Unscoped().Delete(lp).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
