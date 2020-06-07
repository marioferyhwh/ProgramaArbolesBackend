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

//Login funcion de inicio de seccion
func Login(user models.User, m *models.Message) {
	pwd := encriptPasswordUser(user.Password)
	db := configuration.GetConnection()
	defer db.Close()
	db.Where("(nick_name = ? or email = ?) and password = ?", user.Email, user.Email, pwd).First(&user)
	user.Password = ""
	if user.ID <= 0 {
		m.Code = http.StatusUnauthorized
		m.Message = "Verificar Nombre y/o clave"
		return
	}
	// user.Password = ""
	token, err := commons.GenetateJWT(user)
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
func UserCreate(user models.User, m *models.Message) {
	//validacion de datos de usuario
	if user.Email == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta email"
		return
	}
	if user.Actived {
		user.Actived = true
	}
	if user.NickName == "" {
		user.NickName = user.Email
	}
	if user.CodDocumentType == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta tipo de documento"
		return
	}
	if user.Document == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta Documento"
		return
	}
	if user.Password == "" {
		m.Code = http.StatusBadRequest
		m.Message = "falta clave"
		return
	}
	//encriptar clave
	pwd := encriptPasswordUser(user.Password)
	user.Password = pwd
	db := configuration.GetConnection()
	defer db.Close()
	err := createUser(&user, m, db)
	if err != nil {
		return
	}
	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}

//GetUser trae un usuario
func GetUser(user models.User, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getUser(&user, m, db)
	// db.First(&user)
	// err := db.First(&user).Error
	if err != nil {
		m.Code = http.StatusUnauthorized
		m.Message = "no se encontro usuario"
		return
	}
	user.Password = ""
	user.ConfirmPassword = ""
	m.Code = http.StatusOK
	m.Message = "informacion de usuario"
	m.Data = user
}

//EditUser trae un usuario
func EditUser(user models.User, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	db.Begin()
	err := updateUser(&user, m, db)
	m.Code = http.StatusBadRequest
	m.Message = "no se puedo actualizar"
	if err != nil {
		db.Rollback()
		return
	}
	for _, tel := range user.UserTelDelete {
		err = deleteUserTel(&tel, m, db)
		if err != nil {
			db.Rollback()
			return
		}
	}
	for _, tel := range user.UserTelNew {
		err = createUserTel(&tel, m, db)
		if err != nil {
			db.Rollback()
			return
		}
	}
	for _, tel := range user.UserTel {
		err = updateUserTel(&tel, m, db)
		if err != nil {
			db.Rollback()
			return
		}
	}
	user.Password = ""
	user.ConfirmPassword = ""
	db.Commit()
	m.Code = http.StatusOK
	m.Message = "Usuario Editado"
	m.Data = user
}

//DeleteUser trae un usuario
func DeleteUser(user models.User, m *models.Message) {
	//se debe agregar restricciones de borrado
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUser(&user, m, db)
	if err != nil {
		m.Code = http.StatusBadGateway
		m.Message = "no se pudo Borrado Usuario"
		return
	}
	m.Code = http.StatusOK
	m.Message = "Usuario Borrado"
	m.Data = user
}

/*······························································
································································
··············· usuario
································································
······························································*/

//createUser crea usuario con una conexion ya existente
func createUser(u *models.User, m *models.Message, db *gorm.DB) error {
	//	q := `insert into users (created_at,updated_at,actived,nick_name,email,password,cod_document_type,document,name)values(now(),now(),true,$1,$2,$3,$4,$5,$6);`
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	u.Password = ""
	u.ConfirmPassword = ""
	return nil
}

//getUserShort trae usuario con una conexion ya existente
func getUserShort(u *models.User, m *models.Message, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,actived,name").First(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUser trae usuario con una conexion ya existente
func getUser(u *models.User, m *models.Message, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").First(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserList trae usuario con una conexion ya existente
func getUserList(u *[]models.User, m *models.Message, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,actived,name").Find(u).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUser se borra el usuario con una conexion ya existente
func updateUser(u *models.User, m *models.Message, db *gorm.DB) error {
	//q := `update from users set .. where id=?;`
	omitList := []string{"id"}
	if !u.ChangePassword || u.ConfirmPassword == u.Password {
		omitList = append(omitList, "password")
	}
	if !u.ChangeActived {
		omitList = append(omitList, "actived")
	}
	err := db.Model(u).Omit(omitList...).Updates(u).Error
	return err
}

//deleteUser se borra el usuario con una conexion ya existente
func deleteUser(u *models.User, m *models.Message, db *gorm.DB) error {
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

//UserTelCreate crea un nuevo tipo de negocio
func UserTelCreate(ut models.UserTel, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := createUserTel(&ut, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se creo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de negocio creado"
	m.Data = ut
}

//UserTelGet traer un nuevo tipo de negocio
func UserTelGet(ut models.UserTel, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserTel(&ut, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encotro tipo de negocio"
		return
	}
	m.Code = http.StatusOK
	m.Message = "tipo de negocio creado"
	m.Data = ut
}

//UserTelGetList traer lista de tipo de negocio
func UserTelGetList(ut models.UserTel, m *models.Message) {
	uts := []models.UserTel{ut}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserTelList(&uts, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se encontro litado de tipo de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "lista de tipo de negocios"
	m.Data = uts
}

//UserTelUpdate se edita un tipo de negocio
func UserTelUpdate(ut models.UserTel, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUserTel(&ut, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se actualizo"
		return
	}
	m.Code = http.StatusOK
	m.Message = "se actualizo tipo de negocio"
	m.Data = ut
}

//UserTelDelete se borra un tipo de negocio
func UserTelDelete(ut models.UserTel, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUserTel(&ut, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "tipo de negocio no se borro"
		return
	}
	m.Code = http.StatusOK
	m.Message = "borrado correctamente"
	m.Data = ut
}

/*······························································
······························································*/

//createUserTel crea telefono de usuario con una conexion ya existente
func createUserTel(ut *models.UserTel, m *models.Message, db *gorm.DB) error {
	err := db.Create(ut).Error
	return err
}

//getUserTel trae telefono de usuario con una conexion ya existente
func getUserTel(ut *models.UserTel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,phone,descrip").First(ut).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserTelList trae telefono de usuario con una conexion ya existente
func getUserTelList(uts *[]models.UserTel, m *models.Message, db *gorm.DB) error {
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

//updateUserTel se borra el telefono de usuario con una conexion ya existente
func updateUserTel(ut *models.UserTel, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id", "cod_user", "phone"}
	err := db.Model(ut).Omit(omitList...).Updates(ut).Error
	return err
}

//deleteUserTel se borra telefono de usuario con una conexion ya existente
func deleteUserTel(ut *models.UserTel, m *models.Message, db *gorm.DB) error {
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

//createUserLevel crea tipo de documento con una conexion ya existente
func createUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Create(ul).Error
	return err
}

//getUserLevel trae tipo de documento con una conexion ya existente
func getUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,level").First(ul).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//getUserLevelList trae tipo de documento con una conexion ya existente
func getUserLevelList(ul *[]models.UserLevel, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,level").Find(ul).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}

//updateUserLevel se borra el tipo de documento con una conexion ya existente
func updateUserLevel(ul *models.UserLevel, m *models.Message, db *gorm.DB) error {
	omitList := []string{"id"}
	err := db.Model(ul).Omit(omitList...).Updates(ul).Error
	return err
}

//deleteUserLevel se borra el tipo de documento con una conexion ya existente
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
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := createUserCollection(&uc, m, db)
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
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserCollection(&uc, m, db)
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
	ucs := []models.UserCollection{uc}
	db := configuration.GetConnection()
	defer db.Close()
	err := getUserCollectionList(&ucs, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se creo el listado de negocios"
		return
	}
	m.Code = http.StatusOK
	m.Message = "listado de negocios"
	m.Data = ucs
}

//UserCollectionUpdate se edita un enlace entre usuario y cobro
func UserCollectionUpdate(uc models.UserCollection, m *models.Message) {
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUserCollection(&uc, m, db)
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
	if !validateAdmin(m) {
		return
	}
	db := configuration.GetConnection()
	defer db.Close()
	err := deleteUserCollection(&uc, m, db)
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

//createUserCollection crea relacion entre usuario y collection con una conexion ya existente
func createUserCollection(uc *models.UserCollection, m *models.Message, db *gorm.DB) error {
	err := db.Create(uc).Error
	return err
}

//getUserCollection trae  relacion entre usuario y collection con una conexion ya existente
func getUserCollection(uc *models.UserCollection, m *models.Message, db *gorm.DB) error {
	err := db.Select("id,created_at,updated_at,actived,cod_user,cod_user_level,cod_collection,cash,name").First(uc).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//getUserCollectionList trae  relacion entre usuario y collection con una conexion ya existente
func getUserCollectionList(ucs *[]models.UserCollection, m *models.Message, db *gorm.DB) error {
	var uc models.UserCollection
	if len(*ucs) == 1 {
		uc = (*ucs)[0]
	}
	where := ""
	where = fmt.Sprintf("cod_collection = %v", uc.CodCollection)
	if uc.CodUser != 0 {
		where = fmt.Sprintf("cod_user = %v", uc.CodUser)
	}
	fmt.Println("se inicia consulta")
	err := db.Debug().Where(where).Select("id,actived,cod_user,cod_user_level,cod_collection,cash,name").Find(ucs).GetErrors()
	if len(err) != 0 {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUserCollection se borra el  relacion entre usuario y collection con una conexion ya existente
func updateUserCollection(uc *models.UserCollection, m *models.Message, db *gorm.DB) error {
	if uc.ID == 0 {
		return errors.New("no es valido")
	}
	omitList := []string{"id", "cod_collection", "cod_user", ""}
	err := db.Model(uc).Omit(omitList...).Updates(uc).Error
	return err
}

//deleteUserCollection se borra el  relacion entre usuario y collection con una conexion ya existente
func deleteUserCollection(uc *models.UserCollection, m *models.Message, db *gorm.DB) error {
	err := db.Unscoped().Delete(uc).GetErrors()
	if len(err) != 0 {
		return errors.New("Error al borrar")
	}
	return nil
}
