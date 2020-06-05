package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createTelDescrip crea descripcion de los telefonos con una conexion ya existente
func createTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Create(td).Error
	return err
}

//getTelDescrip trae descripcion de los telefonos con una conexion ya existente
func getTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(td).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteTelDescrip se borra el descripcion de los telefonos con una conexion ya existente
func deleteTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(td).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateTelDescrip se borra el descripcion de los telefonos con una conexion ya existente
func updateTelDescrip(td *models.TelDescrip, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(td).Omit(omitList...).Updates(td).Error
	return err
}
