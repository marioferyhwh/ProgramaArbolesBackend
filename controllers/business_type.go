package controllers

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//BusinessTypeCreate crea un nuevo tipo de documento
func BusinessTypeCreate(bt models.BusinessType, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createBusinessType(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de negocio creado"
	m.Data = bt
}

//BusinessTypeGet crea un nuevo tipo de documento
func BusinessTypeGet(bt models.BusinessType, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getBusinessType(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro tipo de negocio"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de negocio creado"
	m.Data = bt
}

//BusinessTypeGetList crea un nuevo tipo de documento
func BusinessTypeGetList(bt models.BusinessType, m *models.Message) {
	bts := []models.BusinessType{bt}
	db := configuration.GetConnection()
	defer db.Close()
	err := getBusinessTypeList(&bts, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se creo el listado de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = bts
}

//BusinessTypeUpdate crea un nuevo tipo de documento
func BusinessTypeUpdate(bt models.BusinessType, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateBusinessType(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo tipo de negocio"
	m.Data = bt
}

//BusinessTypeDelete crea un nuevo tipo de documento
func BusinessTypeDelete(bt models.BusinessType, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteBusinessType(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = bt
}

/*······························································
································································
··············· tipo de negocio
································································
······························································*/

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
