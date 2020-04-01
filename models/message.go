package models

//Message mensaje para el cliente de la API
type Message struct {
	Message string `json:"message"`
	Code int `json:"code"`
}