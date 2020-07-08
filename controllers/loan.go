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
	m.Code = http.StatusBadRequest
	if l.InitialValue <= 0 {
		m.Message = "especifique valor inicial"
		return
	}
	if l.Client.ID <= 0 {
		m.Message = "especifique cliente"
		return
	}
	if l.CodCollection <= 0 {
		m.Message = "especifique cobro"
		return
	}
	if l.CodUser <= 0 {
		m.Message = "especifique usuario"
		return
	}
	if l.Interest <= 0 {
		m.Message = "especifique el interes"
		return
	}
	if l.Quota <= 0 {
		l.Quota = 30
	}
	if l.Balance != l.InitialValue*float32(l.Interest) {
		l.Balance = l.InitialValue * float32(l.Interest)
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := sumCashUserCollection(&models.UserCollection{CodCollection: l.CodCollection, CodUser: l.CodUser}, m, db, -l.InitialValue)
	if err != nil {
		db.Rollback()
		return
	}
	var uc models.Collection
	uc.ID = l.CodCollection
	err = sumBalanceCollection(&uc, m, db, l.Balance)
	if err != nil {
		db.Rollback()
		return
	}
	err = createLoan(&l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "prestamo no se creo"
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "prestamo creado"
	m.Data = l
}

//LoanGet traer un nuevo prestamo
func LoanGet(l models.Loan, m *models.Message) {
	m.Code = http.StatusBadRequest
	if l.ID <= 0 {
		m.Message = "especifique prestamo"
		return
	}
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
	m.Code = http.StatusBadRequest
	if l.CodCollection <= 0 && l.CodClient <= 0 {
		m.Message = "especifique cobro o cliente"
		return
	}
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
	// m.Code = http.StatusBadRequest
	// if l.ID <= 0 {
	// 	m.Message = "especifique prestamo"
	// 	return
	// }
	// db := configuration.GetConnection()
	// defer db.Close()
	// ln := l
	// err := getLoan(&ln, db)
	// if err != nil {
	// 	m.Message = "no se encotro prestamo"
	// 	return
	// }
	// db.Begin()
	// nv := l.Balance - ln.Balance
	// err = sumCashloan(&l, m, db, nv)
	// if err != nil {
	// 	m.Message = "prestamo no se actualizo"
	// 	db.Rollback()
	// 	return
	// }
	// db.Commit()
	m.Code = http.StatusOK
	m.Message = "no se premite actualizar"
	m.Data = l
}

//LoanDelete se borra un prestamo
func LoanDelete(l models.Loan, m *models.Message) {
	m.Code = http.StatusBadRequest
	if l.ID <= 0 {
		m.Message = "especifique prestamo"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := deleteLoan(&l, db)
	if err != nil {
		m.Message = "prestamo no se borro"
		db.Rollback()
		return
	}
	err = sumCashUserCollection(&models.UserCollection{CodCollection: l.CodCollection, CodUser: l.CodUser}, m, db, l.InitialValue)
	if err != nil {
		m.Message = "prestamo no se borro"
		db.Rollback()
		return
	}
	var c models.Collection
	c.ID = l.CodCollection
	err = sumBalanceCollection(&c, m, db, -l.Balance)
	if err != nil {
		m.Message = "prestamo no se borro"
		db.Rollback()
		return
	}
	db.Commit()
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

//getLoan trae Prestamos (id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client,cod_collection,cod_user)
func getLoan(l *models.Loan, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client,cod_collection,cod_user").First(l).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanList trae Prestamos (id,created_at,updated_at,initial_value,interest,quota,balance,cod_loanState,cod_client)
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
	}
	if where != "" && l.CodLoanState != 0 {
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
	omitList := []string{"id", "initial_value", "interest", "quota", "cod_client", "cod_collection", "cod_user", "deleted_at"}
	err := db.Model(l).Omit(omitList...).Save(l).Error
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

//sumCashloan se suma ingresa o se retira un pago
func sumCashloan(l *models.Loan, m *models.Message, db *gorm.DB, nc float32) error {
	var uc models.UserCollection
	uc.CodUser = l.CodUser
	err := getLoan(l, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro Prestamo"
		return err
	}
	l.Balance += nc
	err = updateLoan(l, db)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "no se pudo actualizar"
		return err
	}
	var c models.Collection
	c.ID = l.CodCollection
	err = sumBalanceCollection(&c, m, db, nc)
	if err != nil {
		return err
	}
	uc.CodCollection = l.CodCollection
	err = sumCashUserCollection(&uc, m, db, -nc)
	if err != nil {
		return err
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
	m.Code = http.StatusBadRequest
	if ls.ID <= 0 {
		m.Message = "especifique estado de prestamo"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getLoanState(&ls, db)
	if err != nil {
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
	m.Code = http.StatusBadRequest
	if ls.ID <= 0 {
		m.Message = "especifique estado de prestamo"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateLoanState(&ls, db)
	if err != nil {
		m.Message = "tipo de prestamo no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo estado de prestamo"
	m.Data = ls
}

//LoanStateDelete crea un nuevo estado de prestamo
func LoanStateDelete(ls models.LoanState, m *models.Message) {
	m.Code = http.StatusBadRequest
	if ls.ID <= 0 {
		m.Message = "especifique estado de prestamo"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteLoanState(&ls, db)
	if err != nil {
		m.Message = "estado de prestamo no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "estado de prestamo se borro"
	m.Data = ls
}

/*······························································
······························································*/

//createLoanState crea estado de  prestamos
func createLoanState(ls *models.LoanState, db *gorm.DB) error {
	err := db.Create(ls).Error
	return err
}

//getLoanState trae estado de  prestamos (id,created_at,updated_at,state)
func getLoanState(ls *models.LoanState, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,state").First(ls).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanStateList trae estado de  prestamos (id,state)
func getLoanStateList(lss *[]models.LoanState, db *gorm.DB) error {
	err := db.Select("id,state").Find(lss).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateLoanState se borra el estado de  prestamos
func updateLoanState(ls *models.LoanState, db *gorm.DB) error {
	omitList := []string{"id", "deleted_at"}
	err := db.Model(ls).Omit(omitList...).Save(ls).Error
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
	m.Code = http.StatusBadRequest
	if lp.CodLoan <= 0 {
		m.Message = "falta prestamo"
		return
	}
	if lp.CodCollection <= 0 {
		m.Message = "falta cobro"
		return
	}
	if lp.Cash <= 0 {
		m.Message = "valor no valido"
		return
	}
	if lp.CodUser <= 0 {
		lp.CodUser = m.User.ID
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	var l models.Loan
	l.ID = lp.CodLoan
	l.CodUser = lp.CodUser
	err := sumCashloan(&l, m, db, -lp.Cash)
	if err != nil {
		m.Message = "prestamo no se actualizo"
		db.Rollback()
		return
	}
	err = createLoanPayment(&lp, db)
	if err != nil {
		m.Message = "pago a prestamo no se creo"
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "pago a prestamo creado"
	m.Data = lp
}

//LoanPaymentGet traer un nuevo pago a prestamo
func LoanPaymentGet(lp models.LoanPayment, m *models.Message) {
	m.Code = http.StatusBadRequest
	if lp.ID <= 0 {
		m.Message = "falta identificador de pago de prestamo"
		return
	}
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
	m.Code = http.StatusBadRequest
	if lp.CodLoan <= 0 {
		m.Message = "falta prestamo"
		return
	}
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
	m.Code = http.StatusBadRequest
	if lp.ID <= 0 {
		m.Message = "falta prestamo"
		return
	}
	if lp.Cash < 0 {
		m.Message = "Valor no valido"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	lpn := lp
	err := getLoanPayment(&lpn, db)
	if err != nil {
		m.Message = "pago a prestamo no se encotro"
		db.Rollback()
		return
	}
	db.Begin()
	err = updateLoanPayment(&lp, db)
	if err != nil {
		m.Message = "pago a prestamo no se actualizo"
		db.Rollback()
		return
	}
	var l models.Loan
	l.ID = lp.CodLoan
	err = sumCashloan(&l, m, db, lpn.Cash-lp.Cash)
	if err != nil {
		m.Message = "prestamo no se actualizo"
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "se actualizo pago a prestamo"
	m.Data = lp
}

//LoanPaymentDelete se borra un pago a prestamo
func LoanPaymentDelete(lp models.LoanPayment, m *models.Message) {
	m.Code = http.StatusBadRequest
	if lp.ID <= 0 {
		m.Message = "falta prestamo"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := deleteLoanPayment(&lp, db)
	if err != nil {
		m.Message = "pago a prestamo no se borro"
		db.Rollback()
		return
	}
	var l models.Loan
	l.ID = lp.CodLoan
	err = sumCashloan(&l, m, db, lp.Cash)
	if err != nil {
		m.Message = "prestamo no se actualizo"
		db.Rollback()
		return
	}
	db.Commit()
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

//getLoanPayment trae pagos de prestamos (id,created_at,updated_at,cod_loan,cash,cod_user)
func getLoanPayment(lp *models.LoanPayment, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,cod_loan,cash,cod_user").First(lp).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getLoanPaymentList trae pagos de prestamos (id,updated_at,cod_loan,cash,cod_user)
func getLoanPaymentList(lps *[]models.LoanPayment, db *gorm.DB) error {
	var lp models.LoanPayment
	if len(*lps) == 1 {
		lp = (*lps)[0]
	}
	where := fmt.Sprintf("cod_collection = %v", lp.CodCollection)
	if lp.CodLoan != 0 {
		where = fmt.Sprintf("cod_loan = %v", lp.CodLoan)
	}
	if lp.CodUser != 0 {
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
	omitList := []string{"id", "cod_collection", "cod_user", "cod_loan", "deleted_at"}
	err := db.Model(lp).Omit(omitList...).Save(lp).Error
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
