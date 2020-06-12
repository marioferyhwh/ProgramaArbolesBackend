package controllers

import (
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

/*······························································
································································
··············· varios valores por default
································································
······························································*/

//GeneralValuesGet trae los valores generales de cla pagina
func GeneralValuesGet(m *models.Message) {
	var ge models.GeneralValues
	db := configuration.GetConnection()
	defer db.Close()
	err := getBusinessTypeList(&ge.BusinessTypes, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de tipo de negocios"
		return
	}
	err = getTelDescripList(&ge.TelDescrips, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se creo"
		return
	}
	err = getLoanStateList(&ge.LoanStates, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de pago a prestamos"
		return
	}
	err = getUserLevelList(&ge.UserLevels, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro listado de usuario"
		return
	}
	err = getDocumentTypeList(&ge.DocumentTypes, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de tipo de documentos"
		return
	}
	ge.Levels = commons.UserLevel
	m.Code = http.StatusOK
	m.Message = "lista de tipo de negocios"
	m.Data = ge
}
