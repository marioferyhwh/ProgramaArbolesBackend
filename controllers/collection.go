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
func CollectionCreate(c models.Collection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createCollection(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "cobro creado"
	m.Data = c
}

//CollectionGet traer un nuevo cobro
func CollectionGet(c models.Collection, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollection(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro cobro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "cobro creado"
	m.Data = c
}

//CollectionGetList traer lista de cobro
func CollectionGetList(c models.Collection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	cs := []models.Collection{c}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionList(&cs, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de cobros"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de cobros"
	m.Data = cs
}

//CollectionUpdate se edita un cobro
func CollectionUpdate(c models.Collection, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateCollection(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo cobro"
	m.Data = c
}

//CollectionDelete se borra un cobro
func CollectionDelete(c models.Collection, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteCollection(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cobro no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = c
}

/*······························································
······························································*/

//createCollection crea cobro
func createCollection(c *models.Collection, db *gorm.DB) error {
	err := db.Create(c).Error
	return err
}

//getCollection trae cobro (id,created_at,updated_at,descrip,actived,balance_total)
func getCollection(c *models.Collection, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,descrip,actived,balance_total").First(c).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getCollectionList trae cobro (id,descrip,actived,balance_total)
func getCollectionList(cs *[]models.Collection, db *gorm.DB) error {
	err := db.Select("id,descrip,actived,balance_total").Find(cs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateCollection se borra el cobro
func updateCollection(c *models.Collection, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(c).Omit(omitList...).Updates(c).Error
	return err
}

//deleteCollection se borra el cobro
func deleteCollection(c *models.Collection, db *gorm.DB) error {
	err := db.Unscoped().Delete(c).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//modifiBalanceCollection se suma cash al valor actual
func sumBalanceCollection(c *models.Collection, m *models.Message, db *gorm.DB, nc float32) error {
	err := getCollection(c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro Prestamo"
		return err
	}
	c.BalanceTotal += nc
	err = updateCollection(c, db)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "no se pudo actualizar"
		return err
	}
	return nil
}

/*······························································
································································
··············· lista de movimentos en cobro
································································
······························································*/

//CollectionCashCreate crea un nuevo movimento de cobro
func CollectionCashCreate(cc models.CollectionCash, m *models.Message) {
	m.Code = http.StatusBadRequest
	if cc.CodCollection <= 0 {
		m.Message = "especifique cobro"
		return
	}
	if cc.Cash != 0 {
		m.Message = "valor no valido"
		return
	}
	if cc.CodUser <= 0 {
		cc.CodUser = m.User.ID
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := sumCashUserCollection(&models.UserCollection{CodUser: cc.CodUser, CodCollection: cc.CodCollection}, m, db, cc.Cash)
	if err != nil {
		db.Rollback()
		return
	}
	err = createCollectionCash(&cc, db)
	if err != nil {
		m.Message = "movimento de cobro no se creo"
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "movimento de cobro creado"
	m.Data = cc
}

//CollectionCashGet traer un nuevo movimento de cobro
func CollectionCashGet(cc models.CollectionCash, m *models.Message) {
	if cc.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique movimento"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionCash(&cc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro movimento de cobro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "movimento de cobro creado"
	m.Data = cc
}

//CollectionCashGetList traer lista de movimento de cobro
func CollectionCashGetList(cc models.CollectionCash, m *models.Message) {
	if cc.CodCollection <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	ccs := []models.CollectionCash{cc}
	db := configuration.GetConnection()
	defer db.Close()
	err := getCollectionCashList(&ccs, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no encontro listado el movimento de cobros"
		return
	}
	m.Code = http.StatusOK
	m.Message = "movimento de cobros"
	m.Data = ccs
}

//CollectionCashUpdate se edita un movimento de cobro
func CollectionCashUpdate(cc models.CollectionCash, m *models.Message) {
	m.Code = http.StatusBadRequest
	if cc.ID <= 0 {
		m.Message = "especifique movimento"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	ccn := cc
	err := getCollectionCash(&ccn, db)
	if err != nil {
		m.Message = "no se encontro movimiento"
		return
	}
	db.Begin()
	err = updateCollectionCash(&cc, db)
	if err != nil {
		m.Message = "movimento de cobro no se actualizo"
		db.Rollback()
		return
	}
	err = sumCashUserCollection(&models.UserCollection{CodUser: cc.CodUser, CodCollection: cc.CodCollection}, m, db, (cc.Cash - ccn.Cash))
	if err != nil {
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "se actualizo movimento de cobro"
	m.Data = cc
}

//CollectionCashDelete se borra un movimento de cobro
func CollectionCashDelete(cc models.CollectionCash, m *models.Message) {
	m.Code = http.StatusBadRequest
	if cc.ID <= 0 {
		m.Message = "especifique movimento"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := deleteCollectionCash(&cc, db)
	if err != nil {
		m.Message = "movimento de cobro no se borro"
		db.Rollback()
		return
	}
	err = sumCashUserCollection(&models.UserCollection{CodUser: cc.CodUser, CodCollection: cc.CodCollection}, m, db, -cc.Cash)
	if err != nil {
		db.Rollback()
		return
	}
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = cc
}

/*······························································
······························································*/

//createCollectionCash crea movimento de cobro
func createCollectionCash(cc *models.CollectionCash, db *gorm.DB) error {
	err := db.Create(cc).Error
	return err
}

//getCollectionCash trae movimento de cobro (id,created_at,updated_at,cod_user,cod_collection,cash)
func getCollectionCash(cc *models.CollectionCash, db *gorm.DB) error {
	//.Where("cod_collection == ?", cc.CodCollection)
	err := db.Select("id,created_at,updated_at,cod_user,cod_collection,cash").First(cc).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getCollectionCashList trae lista pagos de cobro (id,cod_user,cod_collection,cash)
func getCollectionCashList(ccs *[]models.CollectionCash, db *gorm.DB) error {
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
func updateCollectionCash(cc *models.CollectionCash, db *gorm.DB) error {
	omitList := []string{"id", "cod_user", "cod_collection"}
	err := db.Model(cc).Omit(omitList...).Updates(cc).Error
	return err
}

//deleteCollectionCash se borra el movimento de cobro
func deleteCollectionCash(cc *models.CollectionCash, db *gorm.DB) error {
	err := db.Unscoped().Delete(cc).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
