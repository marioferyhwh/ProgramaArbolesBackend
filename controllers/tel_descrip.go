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
	if td.Descrip == "" {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion no es valida"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createTelDescrip(&td, db)
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
	if td.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique descripcion"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getTelDescrip(&td, db)
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
	err := getTelDescripList(&tds, db)
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
	if td.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique descripcion"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateTelDescrip(&td, db)
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
	if td.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique descripcion"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getTelDescrip(&td, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se encontro"
		return
	}
	err = deleteTelDescrip(&td, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion del telefono no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion del telefono borro"
	m.Data = td
}

/*······························································
······························································*/

//createTelDescrip crea descripcion de los telefonos
func createTelDescrip(td *models.TelDescrip, db *gorm.DB) error {
	err := db.Create(td).Error
	return err
}

//getTelDescrip trae descripcion de los telefonos (id,created_at,updated_at,descrip)
func getTelDescrip(td *models.TelDescrip, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip").First(td).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getTelDescripList trae descripcion de los telefonos (id,descrip)
func getTelDescripList(tds *[]models.TelDescrip, db *gorm.DB) error {
	err := db.Select("id,descrip").Find(tds).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateTelDescrip se borra el descripcion de los telefonos
func updateTelDescrip(td *models.TelDescrip, db *gorm.DB) error {
	omitList := []string{"id", "deleted_at"}
	err := db.Model(td).Omit(omitList...).Save(td).Error
	return err
}

//deleteTelDescrip se borra el descripcion de los telefonos
func deleteTelDescrip(td *models.TelDescrip, db *gorm.DB) error {
	err := db.Unscoped().Delete(td).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
