package models

//Token permite envolver el token generado de Claim
type Token struct{
	Token string `json:"token"`
}