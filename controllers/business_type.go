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
··············· tipo de negocio
································································
······························································*/

//BusinessTypeCreate crea un nuevo tipo de negocio
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

//BusinessTypeGet traer un nuevo tipo de negocio
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

//BusinessTypeGetList traer lista de tipo de negocio
func BusinessTypeGetList(bt models.BusinessType, m *models.Message) {
	bts := []models.BusinessType{bt}
	db := configuration.GetConnection()
	defer db.Close()
	err := getBusinessTypeList(&bts, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de tipo de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de tipo de negocios"
	m.Data = bts
}

//BusinessTypeUpdate se edita un tipo de negocio
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

//BusinessTypeDelete se borra un tipo de negocio
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
······························································*/

//createBusinessType crea tipo de negocio
func createBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Create(bt).Error
	return err
}

//getBusinessType trae tipo de negocio
func getBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip").First(bt).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getBusinessTypeList trae lista de tipo de negocios
func getBusinessTypeList(bt *[]models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,descrip").Find(bt).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateBusinessType actualizar el tipo de negocio
func updateBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(bt).Omit(omitList...).Updates(bt).Error
	return err
}

//deleteBusinessType borras el tipo de negocio
func deleteBusinessType(bt *models.BusinessType, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(bt).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
