package models

import (
	"github.com/jinzhu/gorm"
)

//User usuario del sistma
type User struct {
	gorm.Model
	/*como va ser una api json tengo que decirle a go que en json el nombre no va en mayuscula 
	sino en miniscula y a gorm le digo que va ser unico y no puede ser nulo*/
	Username string `json:"username" gorm:"not null;unique"`
	Email string    `json:"email"    gorm:"not null;unique"`
	Fullname string `json:"fullname" gorm:"not null"`
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(256)"` //omitalo si esta vacio
	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"` //con gorm:"-" le digo que no lo cree en la base de datos, ni lo busque ni nada
	Picture string  `json:"picture"`
	Comments []Comment `json:"comments,omitempty"`  

}