package  configuration 

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
type configuration struct {
	Server   string
	Port     string
	USer     string
	Password string
	Database string
} 

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal("error:", err, "<-no se pudo leer la configuracion de la base de datos")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal("error:", err, "<-no se pudo decodificar")
	}
	return c
}
//GetConnection conecta a la base de datos
func GetConnection() *gorm.DB {
	c := getConfiguration()
	/* 		"host=myhost port=myport user=gorm dbname=gorm password=mypassword"*/
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Server, c.Port, c.USer, c.Database, c.Password)
	//fmt.Println(dsn)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	
	return db

}
