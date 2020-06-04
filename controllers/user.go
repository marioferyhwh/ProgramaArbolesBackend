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
	err := userCreate(&user, m, db)
	if err != nil {
		return
	}

	m.Code = http.StatusOK
	m.Message = "Usuario Creado"
	m.Data = user
}

//GetUser trae un usario
func GetUser(user models.User, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := getuser(&user, m, db)
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

//EditUser trae un usario
func EditUser(user models.User, m *models.Message) {
	db := configuration.GetConnection()
	defer db.Close()
	err := updateUser(&user, m, db)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "no se puedo actualizar"
		return
	}
	user.Password = ""
	user.ConfirmPassword = ""

	m.Code = http.StatusOK
	m.Message = "Usuario Editado"
	m.Data = user
}

//DeleteUser trae un usario
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

//userCreate crea usuario con una conexion ya existente
func userCreate(u *models.User, m *models.Message, db *gorm.DB) error {
	//	q := `insert into users (created_at,updated_at,actived,nick_name,email,password,cod_document_type,document,name)values(now(),now(),true,$1,$2,$3,$4,$5,$6);`
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	u.Password = ""
	u.ConfirmPassword = ""
	return nil
}

//getuser trae usario con una conexion ya existente
func getuser(u *models.User, m *models.Message, db *gorm.DB) error {
	//q := `select (id,created_at,updated_at,active,nick_name,email,cod_document_type,document,name)from users;`
	err := db.Select("id,created_at,updated_at,actived,nick_name,email,cod_document_type,document,name").Find(u).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//deleteUser se borra el usario con una conexion ya existente
func deleteUser(u *models.User, m *models.Message, db *gorm.DB) error {
	//q := `delete from users where id=?;`
	//q := `update from users set delete_at = now() where id=?;`
	//err :=db.Delete(&u).GetErrors()
	err := db.Unscoped().Delete(u).GetErrors()
	if err != nil {
		return errors.New("no se encuentra")
	}
	return nil
}

//updateUser se borra el usario con una conexion ya existente
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
	if err != nil {
		fmt.Println(err)
		return errors.New("no se encuentra")
	}
	return nil
}
