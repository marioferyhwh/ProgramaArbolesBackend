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
··············· cobro
································································
······························································*/

//CollectionCreate crea un nuevo cobro
func CollectionCreate(bt models.Collection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createCollection(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "cobro creado"
	m.Data = bt
}

//CollectionGet traer un nuevo cobro
func CollectionGet(bt models.Collection, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollection(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro cobro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "cobro creado"
	m.Data = bt
}

//CollectionGetList traer lista de cobro
func CollectionGetList(bt models.Collection, m *models.Message) {
	bts := []models.Collection{bt}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionList(&bts, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de cobros"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de cobros"
	m.Data = bts
}

//CollectionUpdate se edita un cobro
func CollectionUpdate(bt models.Collection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateCollection(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo cobro"
	m.Data = bt
}

//CollectionDelete se borra un cobro
func CollectionDelete(bt models.Collection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteCollection(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = bt
}

/*······························································
······························································*/

//createCollection crea cobro
func createCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Create(c).Error
	return err
}

//getCollection trae cobro
func getCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip,actived,balance_total").First(c).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getCollectionList trae cobro
func getCollectionList(cs *[]models.Collection, m *models.Message, db *gorm.DB) error {
	// var c models.Collection
	// if len(*cs )==0{
	//   c= (*cs)[0]
	// }
	err := db.Select("id,descrip,actived,balance_total").Find(cs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateCollection se borra el cobro
func updateCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(c).Omit(omitList...).Updates(c).Error
	return err
}

//deleteCollection se borra el cobro
func deleteCollection(c *models.Collection, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(c).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· lista de movimentos en cobro
································································
······························································*/

//CollectionCashCreate crea un nuevo movimento de cobro
func CollectionCashCreate(bt models.CollectionCash, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createCollectionCash(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "movimento de cobro no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "movimento de cobro creado"
	m.Data = bt
}

//CollectionCashGet traer un nuevo movimento de cobro
func CollectionCashGet(bt models.CollectionCash, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionCash(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro movimento de cobro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "movimento de cobro creado"
	m.Data = bt
}

//CollectionCashGetList traer lista de movimento de cobro
func CollectionCashGetList(bt models.CollectionCash, m *models.Message) {
	bts := []models.CollectionCash{bt}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionCashList(&bts, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no encontro listado el movimento de cobros"
		return
	}
	m.Code = http.StatusOK
	m.Message = "movimento de cobros"
	m.Data = bts
}

//CollectionCashUpdate se edita un movimento de cobro
func CollectionCashUpdate(bt models.CollectionCash, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateCollectionCash(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "movimento de cobro no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo movimento de cobro"
	m.Data = bt
}

//CollectionCashDelete se borra un movimento de cobro
func CollectionCashDelete(bt models.CollectionCash, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteCollectionCash(&bt, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "movimento de cobro no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = bt
}

/*······························································
······························································*/

//createCollectionCash crea movimento de cobro
func createCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	err := db.Create(cc).Error
	return err
}

//getCollectionCash trae movimento de cobro
func getCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	//.Where("cod_collection == ?", cc.CodCollection)
	err := db.Select("id,created_at,updated_at,cod_user,cod_collection,cash").First(cc).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getCollectionCashList trae lista pagos de cobro
func getCollectionCashList(ccs *[]models.CollectionCash, m *models.Message, db *gorm.DB) error {
	var cc models.CollectionCash
	if len(*ccs) == 1 {
		cc = (*ccs)[0]
	}
	cc.CreatedAt, cc.UpdatedAt = validateTime(models.TimeValidator{I: cc.CreatedAt, E: cc.UpdatedAt})
	err := db.Where("cod_collection == ? and (created_at >= ? and create_at <= ?)", cc.CodCollection, cc.CreatedAt, cc.UpdatedAt).Select("id,cod_user,cod_collection,cash").Find(ccs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateCollectionCash se borra el movimento de cobro
func updateCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "cod_user", "cod_collection"}
	err := db.Model(cc).Omit(omitList...).Updates(cc).Error
	return err
}

//deleteCollectionCash se borra el movimento de cobro
func deleteCollectionCash(cc *models.CollectionCash, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(cc).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
