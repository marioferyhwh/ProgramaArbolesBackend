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
··············· cliente
································································
······························································*/

//ClientCreate crea un nuevo cliente
func ClientCreate(c models.Client, m *models.Message) {
	m.Code = http.StatusBadRequest
	if c.Name == "" {
		m.Message = "falta nombre cliente"
		return
	}
	if c.Adress == "" || c.CodClientListLocation <= 0 {
		m.Message = "falta direccion cliente"
		return
	}
	if c.CodCollection <= 0 {
		m.Message = "falta cobro al que pertenece cliente"
		return
	}
	if c.CodUser <= 0 {
		c.CodUser = m.User.ID
	}
	if c.CodBusinessType <= 0 {
		m.Message = "falta tipo de negocio"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createClient(&c, db)
	m.Code = http.StatusBadRequest
	if err != nil {
		m.Message = "cliente no se creo"
		return
	}
	for _, tel := range c.ClientTel {
		err = createClientTel(&tel, db)
		if err != nil {
			m.Message = "error agregando los telefonos"
			break
		}
	}
	m.Code = http.StatusOK
	m.Message = "cliente creado"
	m.Data = c
}

//ClientGet traer un nuevo cliente
func ClientGet(c models.Client, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClient(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro cliente"
		return
	}
	c.ClientTel = []models.ClientTel{{CodClient: c.ID}}
	err = getClientTelList(&c.ClientTel, db)
	c.Loan = []models.Loan{{CodCollection: c.CodCollection, CodClient: c.ID}}
	err = getLoanList(&c.Loan, db)
	c.User.ID = c.CodUser
	err = getUserShort(&c.User, db)
	m.Code = http.StatusOK
	m.Message = "cliente creado"
	m.Data = c
}

//ClientGetList traer lista de cliente
func ClientGetList(c models.Client, m *models.Message) {
	if c.CodCollection <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	cs := []models.Client{c}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientList(&cs, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se creo el listado de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = cs
}

//ClientUpdate se edita un cliente
func ClientUpdate(c models.Client, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	tx := db.Begin()
	err := updateClient(&c, tx)
	m.Code = http.StatusBadRequest
	m.Message = "no se puedo actualizar"
	if err != nil {
		tx.Rollback()
		return
	}
	for _, tel := range c.ClientTelDelete {
		err = deleteClientTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	for _, tel := range c.ClientTelNew {
		err = createClientTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	for _, tel := range c.ClientTel {
		err = updateClientTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	m.Code = http.StatusOK
	m.Message = "se actualizo cliente"
	m.Data = c
}

//ClientDelete se borra un cliente
func ClientDelete(c models.Client, m *models.Message) {
	if c.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClient(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cliente no se encontro"
		return
	}

	err = deleteClient(&c, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "cliente no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = c
}

/*······························································
······························································*/

//createClient crea cliente
func createClient(c *models.Client, db *gorm.DB) error {
	err := db.Create(c).Error
	return err
}

//getClient trae cliente (id,created_at,updated_at,name,email,cod_document_type,document,Adress,loan_number,cod_collection,cod_loan_state,cod_business_type,cod_client_list_location,cod_user)
func getClient(c *models.Client, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,name,email,cod_document_type,document,Adress,loan_number,cod_collection,cod_loan_state,cod_business_type,cod_client_list_location,cod_user").First(c).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getClientList trae lista de clientes (id,name,loan_number,cod_loan_state)
func getClientList(cs *[]models.Client, db *gorm.DB) error {
	var c models.Client
	if len(*cs) == 1 {
		c = (*cs)[0]
	}
	err := db.Where("cod_collection = ?", c.CodCollection).Select("id,name,loan_number,cod_loan_state").Find(cs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateClient se actualiza el cliente
func updateClient(c *models.Client, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection", "cod_user", "deleted_at"}
	err := db.Model(c).Omit(omitList...).Save(c).Error
	return err
}

//deleteClient se borra el cliente
func deleteClient(c *models.Client, db *gorm.DB) error {
	err := db.Unscoped().Delete(c).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· telefono de clientes
································································
······························································*/

//ClientTelCreate crea un nuevo telefono de cliente
func ClientTelCreate(ct models.ClientTel, m *models.Message) {
	if ct.CodClient <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createClientTel(&ct, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de cliente no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "telefono de cliente creado"
	m.Data = ct
}

//ClientTelGet traer telefono de cliente
func ClientTelGet(ct models.ClientTel, m *models.Message) {
	if ct.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique telefono de cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientTel(&ct, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro telefono de cliente"
		return
	}
	m.Code = http.StatusOK
	m.Message = "telefono de cliente encontrado"
	m.Data = ct
}

//ClientTelGetList traer lista de telefonos de cliente
func ClientTelGetList(ct models.ClientTel, m *models.Message) {
	if ct.CodClient <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cliente"
		return
	}
	cts := []models.ClientTel{ct}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientTelList(&cts, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se creo el telefonos"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = cts
}

//ClientTelUpdate actualiza telefono de cliente
func ClientTelUpdate(ct models.ClientTel, m *models.Message) {
	if ct.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique telefono de cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateClientTel(&ct, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de cliente no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo telefono de cliente"
	m.Data = ct
}

//ClientTelDelete borra telefono de cliente
func ClientTelDelete(ct models.ClientTel, m *models.Message) {
	if ct.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique telefono de cliente"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientTel(&ct, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de cliente no se encontro"
		return
	}
	err = deleteClientTel(&ct, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de cliente no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = ct
}

/*······························································
······························································*/

//createClientTel crea telefonos de clientes
func createClientTel(ct *models.ClientTel, db *gorm.DB) error {
	err := db.Create(ct).Error
	return err
}

//getClientTel trae telefono de cliente (id,created_at,updated_at,phone,cod_tel_descrip)
func getClientTel(ct *models.ClientTel, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,phone,cod_tel_descrip").First(ct).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getClientTelList trae telefonos de cliente (id,phone,cod_tel_descrip)
func getClientTelList(cts *[]models.ClientTel, db *gorm.DB) error {
	var ct models.ClientTel
	if len(*cts) == 1 {
		ct = (*cts)[0]
	}
	err := db.Where("cod_client = ?", ct.CodClient).Select("id,phone,cod_tel_descrip").Find(cts).Limit(100).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateClientTel actualiza telefono de cliente
func updateClientTel(ct *models.ClientTel, db *gorm.DB) error {
	omitList := []string{"id", "cod_client", "deleted_at"}
	err := db.Model(ct).Omit(omitList...).Save(ct).Error
	return err
}

//deleteClientTel se borra el telefono de cliente
func deleteClientTel(ct *models.ClientTel, db *gorm.DB) error {
	err := db.Unscoped().Delete(ct).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· lista de ubicacion de cliente
································································
······························································*/

//ClientLocationCreate crea un nuevo descripcion de ubicacion
func ClientLocationCreate(cll models.ClientListLocation, m *models.Message) {
	if cll.CodCollection <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createClientLocation(&cll, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de ubicacion no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion de ubicacion creado"
	m.Data = cll
}

//ClientLocationGet crea un nuevo descripcion de ubicacion
func ClientLocationGet(cll models.ClientListLocation, m *models.Message) {
	if cll.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique localizacion"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientLocation(&cll, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro descripcion de ubicacion"
		return
	}
	m.Code = http.StatusOK
	m.Message = "descripcion de ubicacion creado"
	m.Data = cll
}

//ClientLocationGetList crea un nuevo descripcion de ubicacion
func ClientLocationGetList(cll models.ClientListLocation, m *models.Message) {
	if cll.CodCollection <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique cobro"
		return
	}
	clls := []models.ClientListLocation{cll}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientLocationList(&clls, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se creo el listado de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = clls
}

//ClientLocationUpdate crea un nuevo descripcion de ubicacion
func ClientLocationUpdate(cll models.ClientListLocation, m *models.Message) {
	if cll.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique localizacion"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateClientLocation(&cll, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de ubicacion no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo descripcion de ubicacion"
	m.Data = cll
}

//ClientLocationDelete crea un nuevo descripcion de ubicacion
func ClientLocationDelete(cll models.ClientListLocation, m *models.Message) {
	if cll.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique localizacion"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getClientLocation(&cll, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de ubicacion no se encontro"
		return
	}
	err = deleteClientLocation(&cll, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "descripcion de ubicacion no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = cll
}

/*······························································
······························································*/

//createClientLocation crea ubicacion valida para clientes por collection
func createClientLocation(cll *models.ClientListLocation, db *gorm.DB) error {
	err := db.Create(cll).Error
	return err
}

//getClientLocation trae ubicacion valida para clientes por collection (id,created_at,updated_at,descrip)
func getClientLocation(cll *models.ClientListLocation, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,cod_collection,descrip").First(cll).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getClientLocationList trae lista de ubicacion valida para clientes por collection (id,descrip)
func getClientLocationList(clls *[]models.ClientListLocation, db *gorm.DB) error {
	var cll models.ClientListLocation
	if len(*clls) == 1 {
		cll = (*clls)[0]
	}
	err := db.Where("cod_collection = ?", cll.CodCollection).Select("id,descrip").Find(clls).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateClientLocation actualizar el ubicacion valida para clientes por collection
func updateClientLocation(cll *models.ClientListLocation, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection", "deleted_at"}
	err := db.Model(cll).Omit(omitList...).Save(cll).Error
	return err
}

//deleteClientLocation borra el ubicacion valida para clientes por collection
func deleteClientLocation(cll *models.ClientListLocation, db *gorm.DB) error {
	err := db.Unscoped().Delete(cll).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
