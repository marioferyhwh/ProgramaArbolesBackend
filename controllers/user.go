package controllers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/configuration"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

type model struct {
	Email    string
	Password string
}

//encriptPasswordUser Se encripta la clave
func encriptPasswordUser(password string) string {
	c := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", c)
}

/*······························································
································································
··············· usuario
································································
······························································*/

//Login funcion de inicio de seccion
func Login(u models.User, m *models.Message) {
	pwd := encriptPasswordUser(u.Password)
	db := configuration.GetConnection()
	defer db.Close()
	// db.Debug().Where("(nick_name = ? or email = ?) and password = ? and actived = true", u.Email, u.Email, pwd).First(&u)
	db.Where("(email = ?) and password = ? and actived = true", u.Email, pwd).First(&u)
	u.Password = ""
	if u.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y/o clave"
		return
	}
	// u.Password = ""
	token, err := commons.GenetateJWT(u)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "error generando token"
		return
	}
	/*
		fmt.Print(token)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = err.Error()
		return
		}
	*/
	m.Code = http.StatusOK
	m.NewToken = token
}

//UserCreate crea el usuario
func UserCreate(u models.User, m *models.Message) {
	m.Code = http.StatusBadRequest
	if u.Email == "" {
		m.Message = "falta email"
		return
	}
	if u.CodDocumentType == "" {
		m.Message = "falta tipo de documento"
		return
	}
	if u.Document == "" {
		m.Message = "falta Documento"
		return
	}
	if u.Password == "" {
		m.Message = "falta clave"
		return
	}
	if u.Actived {
		u.Actived = true
	}
	if u.NickName == "" {
		u.NickName = u.Email
	}
	//encriptar clave
	pwd := encriptPasswordUser(u.Password)
	u.Password = pwd
	db := configuration.GetConnection()
	defer db.Close()
	err := createUser(&u, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "user no se creo"
		return
	}
	for _, tel := range u.UserTel {
		err = createUserTel(&tel, db)
		if err != nil {
			break
		}
	}
	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = u
}

//UserGet trae un usuario
func UserGet(u models.User, m *models.Message) {
	if u.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique usuario"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUser(&u, db)
	// db.First(&u)
	// err := db.First(&u).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro usuario"
		return
	}
	u.Password = ""
	u.ConfirmPassword = ""
	m.Code = http.StatusOK
	m.Message = "informacion de usuario"
	m.Data = u
}

//UserGetList pendiente por desarrollar
func UserGetList(u models.User, m *models.Message) {
	if u.ID != 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique usuario"
		return
	}
	var us = []models.User{u}

	db := configuration.GetConnection()
	defer db.Close()
	err := getUserList(&us, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro usuario"
		return
	}
	m.Code = http.StatusOK
	m.Message = "informacion de usuario"
	m.Data = us
}

//UserUpdate trae un usuario
func UserUpdate(u models.User, m *models.Message) {
	if u.ID <= 0 {
		m.Code = http.StatusBadRequest
		m.Message = "especifique usuario"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	var err error
	m.Code = http.StatusBadRequest
	m.Message = "no se puedo actualizar"
	tx := db.Begin()
	for _, tel := range u.UserTelDelete {
		err = deleteUserTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	for _, tel := range u.UserTelNew {
		err = createUserTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	for _, tel := range u.UserTel {
		err = updateUserTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = updateUser(&u, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	u.Password = ""
	u.ConfirmPassword = ""
	tx.Commit()
	m.Code = http.StatusOK
	m.Message = "Usuario Editado"
	m.Data = u
}

//UserDelete trae un usuario
func UserDelete(u models.User, m *models.Message) {
	m.Code = http.StatusBadRequest
	if u.ID <= 0 {
		m.Message = "especifique usuario"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	tx := db.Begin()
	var err error
	m.Message = "no se pudo Borrado Usuario"
	for _, tel := range u.UserTel {
		err = deleteUserTel(&tel, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = deleteUser(&u, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	m.Code = http.StatusOK
	m.Message = "Usuario Borrado"
	m.Data = u
}

/*······························································
······························································*/

//createUser crea usuario
func createUser(u *models.User, db *gorm.DB) error {
	//	q := `insert into users (created_at,updated_at,actived,nick_name,email,password,cod_document_type,document,name)values(now(),now(),true,$1,$2,$3,$4,$5,$6);`
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	u.Password = ""
	u.ConfirmPassword = ""
	return nil
}

//getUserShort trae usuario (id,actived,name)
func getUserShort(u *models.User, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,actived,name").First(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUser trae usuario (id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name)
func getUser(u *models.User, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name,time_zone").First(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserList trae usuario (id,actived,name)
func getUserList(u *[]models.User, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,actived,name").Find(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUser se borra el usuario
func updateUser(u *models.User, db *gorm.DB) error {
	//q := `update from users set .. where id=?;`
	omitList := []string{"id", "deleted_at"} //"CodDocumentType", "Document"
	if !u.ChangePassword || u.ConfirmPassword == u.Password {
		omitList = append(omitList, "password", "passeword_at")
	}
	if !u.ChangeActived {
		omitList = append(omitList, "actived")
	}
	err := db.Model(u).Omit(omitList...).Save(u).Error
	return err
}

//deleteUser se borra el usuario
func deleteUser(u *models.User, db *gorm.DB) error {
	//q := `delete from users where id=?;`
	//q := `update from users set delete_at = now() where id=?;`
	//err :=db.Delete(&u).GetErrors()
	err := db.Unscoped().Delete(u).GetErrors()
	if len(err) != 0 {
		return errors.New("Erro al borrar")
	}
	return nil
}

/*······························································
································································
··············· telefono del usuario
································································
······························································*/

//UserTelCreate crea un nuevo telefono de usuario
func UserTelCreate(ut models.UserTel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ut.CodUser <= 0 {
		m.Message = "no se especifico a que usuario"
		return
	}
	if ut.Phone == "" {
		m.Message = "introdusca numero telefonico"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createUserTel(&ut, db)
	if err != nil {
		m.Message = "telefono de usuario no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "telefono de usuario creado"
	m.Data = ut
}

//UserTelGet traer un nuevo telefono de usuario
func UserTelGet(ut models.UserTel, m *models.Message) {
	m.Code = http.StatusBadRequest
	if ut.ID <= 0 {
		m.Message = "especifique telefono"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserTel(&ut, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro telefono de usuario"
		return
	}
	m.Code = http.StatusOK
	m.Message = "telefono de usuario creado"
	m.Data = ut
}

//UserTelGetList traer lista de telefono de usuario
func UserTelGetList(ut models.UserTel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ut.CodUser <= 0 {
		m.Message = "no se especifico a que usuario"
		return
	}
	uts := []models.UserTel{ut}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserTelList(&uts, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de telefono de usuarios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de telefono de usuarios"
	m.Data = uts
}

//UserTelUpdate se edita un telefono de usuario
func UserTelUpdate(ut models.UserTel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ut.ID <= 0 {
		m.Message = "especifique numero"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUserTel(&ut, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de usuario no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo telefono de usuario"
	m.Data = ut
}

//UserTelDelete se borra un telefono de usuario
func UserTelDelete(ut models.UserTel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ut.ID <= 0 {
		m.Message = "especifique numero"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUserTel(&ut, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "telefono de usuario no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = ut
}

/*······························································
······························································*/

//createUserTel crea telefono de usuario
func createUserTel(ut *models.UserTel, db *gorm.DB) error {
	err := db.Create(ut).Error
	return err
}

//getUserTel trae telefono de usuario (id,created_at,updated_at,phone,descrip)
func getUserTel(ut *models.UserTel, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,phone,descrip").First(ut).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserTelList trae telefono de usuario (id,phone,descrip)
func getUserTelList(uts *[]models.UserTel, db *gorm.DB) error {
	var ut models.UserTel
	if len(*uts) == 1 {
		ut = (*uts)[0]
	}
	err := db.Where("cod_user = ?", ut.CodUser).Select("id,phone,descrip").Find(uts).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUserTel se borra el telefono de usuario
func updateUserTel(ut *models.UserTel, db *gorm.DB) error {
	omitList := []string{"id", "cod_user", "deleted_at"}
	err := db.Model(ut).Omit(omitList...).Save(ut).Error
	return err
}

//deleteUserTel se borra telefono de usuario
func deleteUserTel(ut *models.UserTel, db *gorm.DB) error {
	err := db.Unscoped().Delete(ut).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· Nivel de usuario
································································
······························································*/

//UserLevelCreate crea un nuevo tipo de documento
func UserLevelCreate(ul models.UserLevel, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createUserLevel(&ul, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de usuario no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de usuario creado"
	m.Data = ul
}

//UserLevelGet crea un nuevo tipo de documento
func UserLevelGet(ul models.UserLevel, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	m.Code = http.StatusBadGateway
	if ul.ID <= 0 {
		m.Message = "especifique nivel de usuario"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserLevel(&ul, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro tipo de usuario"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de usuario creado"
	m.Data = ul
}

//UserLevelGetList crea un nuevo tipo de documento
func UserLevelGetList(ul models.UserLevel, m *models.Message) {
	uls := []models.UserLevel{ul}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserLevelList(&uls, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro listado de usuario"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de usuario"
	m.Data = uls
}

//UserLevelUpdate crea un nuevo tipo de documento
func UserLevelUpdate(ul models.UserLevel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ul.ID <= 0 {
		m.Message = "especifique nivel de usuario"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUserLevel(&ul, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de usuario no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo tipo de usuario"
	m.Data = ul
}

//UserLevelDelete crea un nuevo tipo de documento
func UserLevelDelete(ul models.UserLevel, m *models.Message) {
	m.Code = http.StatusBadGateway
	if ul.ID <= 0 {
		m.Message = "especifique nivel de usuario"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUserLevel(&ul, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de usuario no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = ul
}

/*······························································
······························································*/

//createUserLevel crea nivel de usuario
func createUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Create(ul).Error
	return err
}

//getUserLevel trae nivel de usuario (id,created_at,updated_at,level)
func getUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,level").First(ul).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//getUserLevelList trae nivel de usuario (id,level)
func getUserLevelList(ul *[]models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,level").Find(ul).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateUserLevel se borra el nivel de usuario
func updateUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "deleted_at"}
	err := db.Model(ul).Omit(omitList...).Save(ul).Error
	return err
}

//deleteUserLevel se borra el nivel de usuario
func deleteUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(ul).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

/*······························································
································································
··············· lista de collection por usuario
································································
······························································*/

//UserCollectionCreate crea un nuevo enlace entre usuario y cobro
func UserCollectionCreate(uc models.UserCollection, m *models.Message) {
	m.Code = http.StatusBadGateway
	if uc.CodCollection <= 0 {
		m.Message = "especifique cobro"
		return
	}
	if uc.CodUser <= 0 {
		m.Message = "especifique usuario"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createUserCollection(&uc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "enlace entre usuario y cobro no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "enlace entre usuario y cobro creado"
	m.Data = uc
}

//UserCollectionGet traer un nuevo enlace entre usuario y cobro
func UserCollectionGet(uc models.UserCollection, m *models.Message) {
	m.Code = http.StatusBadGateway
	if uc.ID <= 0 && (uc.CodCollection <= 0 || uc.CodUser <= 0) {
		m.Message = "especifique cobro"
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserCollection(&uc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro enlace entre usuario y cobro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "enlace entre usuario y cobro creado"
	m.Data = uc
}

//UserCollectionGetList crea un nuevo tipo de documento
func UserCollectionGetList(uc models.UserCollection, m *models.Message) {
	m.Code = http.StatusBadGateway
	if uc.CodCollection <= 0 && uc.CodUser <= 0 {
		m.Message = "especifique cobro y/o usuario"
		return
	}
	ucs := []models.UserCollection{uc}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserCollectionList(&ucs, db)
	if err != nil {
		m.Message = "no se creo el listado de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = ucs
}

//UserCollectionUpdate se edita un enlace entre usuario y cobro
func UserCollectionUpdate(uc models.UserCollection, m *models.Message) {
	m.Code = http.StatusBadGateway
	if uc.ID <= 0 {
		m.Message = "especifique cobro"
		return
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUserCollection(&uc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "enlace entre usuario y cobro no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo enlace entre usuario y cobro"
	m.Data = uc
}

//UserCollectionDelete se borra un enlace entre usuario y cobro
func UserCollectionDelete(uc models.UserCollection, m *models.Message) {
	m.Code = http.StatusBadGateway
	if uc.ID <= 0 {
		m.Message = "especifique cobro"
	}
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUserCollection(&uc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "enlace entre usuario y cobro no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = uc
}

/*······························································
······························································*/

//createUserCollection crea relacion entre usuario y collection
func createUserCollection(uc *models.UserCollection, db *gorm.DB) error {
	err := db.Create(uc).Error
	return err
}

//getUserCollection trae  relacion entre usuario y collection (id,created_at,updated_at,actived,cod_user,cod_user_level,cod_collection,cash,name)
func getUserCollection(uc *models.UserCollection, db *gorm.DB) error {
	where := ""
	if uc.ID <= 0 && (uc.CodCollection > 0 && uc.CodUser > 0) {
		where = fmt.Sprintf("cod_collection = %v and cod_user =%v", uc.CodCollection, uc.CodUser)
	}
	err := db.Where(where).Select("id,created_at,updated_at,actived,cod_user,cod_user_level,cod_collection,cash,name").First(uc).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserCollectionList trae  relacion entre usuario y collection (id,actived,cod_user,cod_user_level,cod_collection,cash,name)
func getUserCollectionList(ucs *[]models.UserCollection, db *gorm.DB) error {
	var uc models.UserCollection
	if len(*ucs) == 1 {
		uc = (*ucs)[0]
	}
	where := ""
	where = fmt.Sprintf("cod_collection = %v", uc.CodCollection)
	if uc.CodUser != 0 {
		where = fmt.Sprintf("cod_user = %v", uc.CodUser)
	}
	err := db.Where(where).Select("id,actived,cod_user,cod_user_level,cod_collection,cash,name").Find(ucs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUserCollection se borra el  relacion entre usuario y collection
func updateUserCollection(uc *models.UserCollection, db *gorm.DB) error {
	omitList := []string{"id", "cod_collection", "cod_user", "deleted_at"}
	err := db.Model(uc).Omit(omitList...).Save(uc).Error
	return err
}

//deleteUserCollection se borra el  relacion entre usuario y collection
func deleteUserCollection(uc *models.UserCollection, db *gorm.DB) error {
	err := db.Unscoped().Delete(uc).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//modifiCashUserCollection se suma cash al valor actual
func sumCashUserCollection(uc *models.UserCollection, m *models.Message, db *gorm.DB, nc float32) error {
	err := getUserCollection(uc, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro relacion"
		return err
	}
	uc.Cash = uc.Cash + nc
	err = updateUserCollection(uc, db)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "no se pudo actualizar"
		return err
	}
	return nil
}
