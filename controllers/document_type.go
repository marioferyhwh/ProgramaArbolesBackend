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
···············  tipos de documentos
································································
······························································*/

//DocumentTypeCreate crea un nuevo tipo de documento
func DocumentTypeCreate(dt models.DocumentType, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createDocumentType(&dt, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "documento no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "documento creado"
	m.Data = dt
}

//DocumentTypeGet traer un nuevo tipo de documento
func DocumentTypeGet(dt models.DocumentType, m *models.Message) {
	if dt.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique typo de documento"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getDocumentType(&dt, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro tipo de documento"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de documento creado"
	m.Data = dt
}

//DocumentTypeGetList traer lista de tipo de documento
func DocumentTypeGetList(dt models.DocumentType, m *models.Message) {
	dts := []models.DocumentType{dt}
	db := configuration.GetConnection()
	defer db.Close()
	err := getDocumentTypeList(&dts, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de tipo de documentos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de tipo de documentos"
	m.Data = dts
}

//DocumentTypeUpdate se edita un tipo de documento
func DocumentTypeUpdate(dt models.DocumentType, m *models.Message) {
	if dt.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique typo de documento"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateDocumentType(&dt, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de documento no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo tipo de documento"
	m.Data = dt
}

//DocumentTypeDelete crea un nuevo tipo de documento
func DocumentTypeDelete(dt models.DocumentType, m *models.Message) {
	if dt.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique typo de documento"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getDocumentType(&dt, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "documento no se encontro"
		return
	}
	err = deleteDocumentType(&dt, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "documento no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "documento se borro"
	m.Data = dt
}

/*······························································
······························································*/

//createDocumentType crea tipo de documento
func createDocumentType(dt *models.DocumentType, db *gorm.DB) error {
	err := db.Create(dt).Error
	return err
}

//getDocumentType trae tipo de documento (id,created_at,updated_at,name_short,descrip)
func getDocumentType(dt *models.DocumentType, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,name_short,descrip").First(dt).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getDocumentTypeList trae tipo de documento (id,name_short,descrip)
func getDocumentTypeList(dts *[]models.DocumentType, db *gorm.DB) error {
	err := db.Select("id,name_short,descrip").Find(dts).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateDocumentType se borra el tipo de documento
func updateDocumentType(dt *models.DocumentType, db *gorm.DB) error {
	omitList := []string{"id", "name_short", "deleted_at"}
	err := db.Model(dt).Omit(omitList...).Save(dt).Error
	return err
}

//deleteDocumentType se borra el tipo de documento
func deleteDocumentType(dt *models.DocumentType, db *gorm.DB) error {
	err := db.Unscoped().Delete(dt).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
