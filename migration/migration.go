package migration

//este archivo se ejecutará spolo una vez, será el que me cree las tablas en la base de datos

import (
	"github.com/EnriqueLeonCameron/edcomments/configuration"
	"github.com/EnriqueLeonCameron/edcomments/models"
)

func Migrate()  {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	//voy a unir los campos UserID y CommentID para hacer una llave unica con ellos y que un 
	//usuario pueda votar solo una vez en un mismo comentario
	                                 //nombre del indice unico  //nom campos que van a armar el indice
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id","user_id")
}