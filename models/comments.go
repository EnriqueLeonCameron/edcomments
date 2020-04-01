package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
    gorm.Model //implicito -> id, fecha creacion, fecha modificacion, etc
	UserID uint        `json:"userId"`
	ParentID uint      `json:"parentId"` //id del commentario padre
	Votes uint32       `json:"votes"`
	Content string     `json:"content"`
	HasVote int8       `json:"hasVote" gorm:"-"`//para que el usuraio solo pueda votar positiva o negativa mente una vez, que no pueda poner 3 votos positivos al mismo comentario por ejemplo
	User []User        `json:"user,omitempty"` //le digo que tendre un slice de usuarios pero en realidad solo será un usuario, el que hizo el commentario, en este slice tendré toda la info del usuario que comentó
	Children []Comment `json:"childre,omitempty"`//todos los comentarios hijos o comentarios respuesta
}