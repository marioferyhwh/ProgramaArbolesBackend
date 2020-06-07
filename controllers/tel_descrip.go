package controllers

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

/*······························································
································································
··············· descripcion de telefonos
································································
······························································*/

//TelDescripCreate crea un nuevo tipo de telefono
func TelDescripCreate(td models.TelDescrip, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createTelDescrip(&td, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion del telefono creado"
	m.Data = td
}

//TelDescripGet traer un nuevo descripcion de telefono
func TelDescripGet(td models.TelDescrip, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getTelDescrip(&td, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro descripcion de telefono"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion de telefono creado"
	m.Data = td
}

//TelDescripGetList crea un nuevo tipo de telefono
func TelDescripGetList(td models.TelDescrip, m *models.Message) {
	tds := []models.TelDescrip{td}
	db := configuration.GetConnection()
	defer db.Close()
	err := getTelDescripList(&tds, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion del telefono borro"
	m.Data = tds
}

//TelDescripUpdate se edita un descripcion de telefono
func TelDescripUpdate(td models.TelDescrip, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateTelDescrip(&td, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de telefono no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo descripcion de telefono"
	m.Data = td
}

//TelDescripDelete crea un nuevo tipo de telefono
func TelDescripDelete(td models.TelDescrip, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteTelDescrip(&td, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion del telefono borro"
	m.Data = td
}

/*······························································
······························································*/

//createTelDescrip crea descripcion de los telefonos con una conexion ya existente
func createTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Create(td).Error
	return err
}

//getTelDescrip trae descripcion de los telefonos con una conexion ya existente
func getTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip").First(td).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getTelDescripList trae descripcion de los telefonos con una conexion ya existente
func getTelDescripList(tds *[]models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,descrip").Find(tds).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateTelDescrip se borra el descripcion de los telefonos con una conexion ya existente
func updateTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(td).Omit(omitList...).Updates(td).Error
	return err
}

//deleteTelDescrip se borra el descripcion de los telefonos con una conexion ya existente
func deleteTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(td).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
