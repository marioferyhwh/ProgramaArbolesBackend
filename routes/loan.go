package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//SetLoanCreateRoutes Creacion de prestamo
func SetLoanCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var l models.Loan
	err := c.Bind(&l)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.LoanCreate(l, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanGetRoutes traer prestamo
func SetLoanGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var l models.Loan
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	l.ID = id
	controllers.LoanGet(l, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanGetListRoutes listado de ubicaiones
func SetLoanGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var l models.Loan
	id, err := getid32(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	l.CodCollection = id
	controllers.LoanGetList(l, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanEditRoutes actualizar prestamo
func SetLoanEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var l models.Loan
	err := c.Bind(&l)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	l.ID = id
	controllers.LoanUpdate(l, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanDeleteRoutes borrar prestamo
func SetLoanDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var l models.Loan
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	l.ID = id
	controllers.LoanDelete(l, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· estados de prestamos
································································
······························································*/

//SetLoanStateCreateRoutes Creacion de estado de prestamo
func SetLoanStateCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ls models.LoanState
	err := c.Bind(&ls)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego estado de prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.LoanStateCreate(ls, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanStateGetRoutes traer estado de prestamo
func SetLoanStateGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ls models.LoanState
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de estado de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	ls.ID = uint8(id)
	controllers.LoanStateGet(ls, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanStateGetListRoutes listado de ubicaiones
func SetLoanStateGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ls models.LoanState
	controllers.LoanStateGetList(ls, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanStateEditRoutes actualizar estado de prestamo
func SetLoanStateEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ls models.LoanState
	err := c.Bind(&ls)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego estado de prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de estado de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	ls.ID = uint8(id)
	controllers.LoanStateUpdate(ls, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanStateDeleteRoutes borrar estado de prestamo
func SetLoanStateDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var ls models.LoanState
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de estado de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	ls.ID = uint8(id)
	controllers.LoanStateDelete(ls, &m)
	return commons.DisplayMessage(c, &m)
}

/*······························································
································································
··············· pago de prestamo
································································
······························································*/

//SetLoanPaymentCreateRoutes Creacion de pago de prestamo
func SetLoanPaymentCreateRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var lp models.LoanPayment
	err := c.Bind(&lp)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego pago de prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	controllers.LoanPaymentCreate(lp, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanPaymentGetRoutes traer pago de prestamo
func SetLoanPaymentGetRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var lp models.LoanPayment
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de pago de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	lp.ID = id
	controllers.LoanPaymentGet(lp, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanPaymentGetListRoutes listado de ubicaiones
func SetLoanPaymentGetListRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var lp models.LoanPayment
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de pago de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	lp.CodLoan = id
	controllers.LoanPaymentGetList(lp, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanPaymentEditRoutes actualizar pago de prestamo
func SetLoanPaymentEditRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var lp models.LoanPayment
	err := c.Bind(&lp)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprint("no llego pago de prestamo ->", err)
		return commons.DisplayMessage(c, &m)
	}
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de pago de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	lp.ID = id
	controllers.LoanPaymentUpdate(lp, &m)
	return commons.DisplayMessage(c, &m)
}

//SetLoanPaymentDeleteRoutes borrar pago de prestamo
func SetLoanPaymentDeleteRoutes(c echo.Context) error {
	var m models.Message
	getUserInterface(c, &m.User)
	var lp models.LoanPayment
	id, err := getid(c)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "identificador de pago de prestamo no valido"
		return commons.DisplayMessage(c, &m)
	}
	lp.ID = id
	controllers.LoanPaymentDelete(lp, &m)
	return commons.DisplayMessage(c, &m)
}
