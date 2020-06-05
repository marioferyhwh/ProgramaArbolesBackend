package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createDocumentType crea tipo de documento con una conexion ya existente
func createDocumentType(dt *models.DocumentType, m *models.Message, db *gorm.DB) error {
	err := db.Create(dt).Error
	return err
}

//getDocumentType trae tipo de documento con una conexion ya existente
func getDocumentType(dt *models.DocumentType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,name_short,descrip").First(dt).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//getDocumentTypeList trae tipo de documento con una conexion ya existente
func getDocumentTypeList(dt *[]models.DocumentType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,name_short,descrip").Find(dt).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateDocumentType se borra el tipo de documento con una conexion ya existente
func updateDocumentType(dt *models.DocumentType, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(dt).Omit(omitList...).Updates(dt).Error
	return err
}

//deleteDocumentType se borra el tipo de documento con una conexion ya existente
func deleteDocumentType(dt *models.DocumentType, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(dt).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}
