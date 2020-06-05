package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createClient crea clientes con una conexion ya existente
func createClient(c *models.Client, m *models.Message, db *gorm.DB) error {
	err := db.Create(c).Error
	return err
}

//getClient trae clientes con una conexion ya existente
func getClient(c *models.Client, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(c).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteClient se borra el clientes con una conexion ya existente
func deleteClient(c *models.Client, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(c).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateClient se borra el clientes con una conexion ya existente
func updateClient(c *models.Client, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(c).Omit(omitList...).Updates(c).Error
	return err
}

/*······························································
································································
··············· telefono de clientes
································································
······························································*/

//createClientTel crea telefonos de clientes con una conexion ya existente
func createClientTel(ct *models.ClientTel, m *models.Message, db *gorm.DB) error {
	err := db.Create(ct).Error
	return err
}

//getClientTel trae telefonos de clientes con una conexion ya existente
func getClientTel(ct *models.ClientTel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(ct).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteClientTel se borra el telefonos de clientes con una conexion ya existente
func deleteClientTel(ct *models.ClientTel, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ct).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateClientTel se borra el telefonos de clientes con una conexion ya existente
func updateClientTel(ct *models.ClientTel, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ct).Omit(omitList...).Updates(ct).Error
	return err
}

/*······························································
································································
··············· lista de ubicaion de cliente
································································
······························································*/

//createClientListLocation crea ubicacion valida para clientes por collection con una conexion ya existente
func createClientListLocation(cll *models.ClientListLocation, m *models.Message, db *gorm.DB) error {
	err := db.Create(cll).Error
	return err
}

//getClientListLocation trae ubicacion valida para clientes por collection con una conexion ya existente
func getClientListLocation(cll *models.ClientListLocation, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(cll).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteClientListLocation se borra el ubicacion valida para clientes por collection con una conexion ya existente
func deleteClientListLocation(cll *models.ClientListLocation, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(cll).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateClientListLocation se borra el ubicacion valida para clientes por collection con una conexion ya existente
func updateClientListLocation(cll *models.ClientListLocation, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(cll).Omit(omitList...).Updates(cll).Error
	return err
}
