package models

import jwt "github.com/dgrijalva/jwt-go"
//jwt es un alias para github.com/dgrijalva/jwt-go, para no tener que estar escribiendo jwt-go

//Claim Token de usuario
type Claim struct {
	User `json:"user"` //solo lo escribo una vez por que se que este campo va ser de tipo user por que ya tengo una estructura quie se llama asi, es lo mis mo que -> User User, se llama user de tipo user
	jwt.StandardClaims //implicito, fecha de expirac ion etc
}