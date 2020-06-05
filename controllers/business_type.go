package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createBusinessType crea tipo de negocios con una conexion ya existente
func createBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Create(bt).Error
	return err
}

//getBusinessType trae tipo de negocios con una conexion ya existente
func getBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip").First(bt).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//getBusinessTypeList trae tipo de negocios con una conexion ya existente
func getBusinessTypeList(bt *[]models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,descrip").Find(bt).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateBusinessType se borra el tipo de negocios con una conexion ya existente
func updateBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(bt).Omit(omitList...).Updates(bt).Error
	return err
}

//deleteBusinessType se borra el tipo de negocios con una conexion ya existente
func deleteBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(bt).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}
