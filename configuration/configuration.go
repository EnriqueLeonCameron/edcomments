package configuration

import (
	"fmt"
	"os"
	"log"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)//aqui toma la info del archivo y la mapeo a la estructura c
	if err != nil {
		log.Fatal(err)
	}

	return c
}

//funcion que devuelve una conexion a la base de datos
func GetConnection() *gorm.DB {
	c := GetConfiguration() //la configuracion de la Base de datos //le doy la info del archivo json
	//user:password@tcp(server:port)/nombreDataBase?charset=utf8&parseTime=true&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	//Sprintf me devuelve un string formateado con la info que yo le pase
	//%s -> argumento de tipo string
 
	db, err := gorm.Open("mysql",dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}