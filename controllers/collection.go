package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//createCollection crea collection con una conexion ya existente
func createCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Create(c).Error
	return err
}

//getCollection trae collection con una conexion ya existente
func getCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(c).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteCollection se borra el collection con una conexion ya existente
func deleteCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(c).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateCollection se borra el collection con una conexion ya existente
func updateCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(c).Omit(omitList...).Updates(c).Error
	return err
}

/*······························································
································································
··············· lista de movimentos en collection
································································
······························································*/

//createCollectionCash crea movimento de collection con una conexion ya existente
func createCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	err := db.Create(cc).Error
	return err
}

//getCollectionCash trae movimento de collection con una conexion ya existente
func getCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(cc).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteCollectionCash se borra el movimento de collection con una conexion ya existente
func deleteCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(cc).GetErrors()
	if err != nil {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateCollectionCash se borra el movimento de collection con una conexion ya existente
func updateCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(cc).Omit(omitList...).Updates(cc).Error
	return err
}
