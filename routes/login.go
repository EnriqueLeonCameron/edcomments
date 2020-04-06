package routes

import (
	"github.com/gorilla/mux"

	"github.com/EnriqueLeonCameron/edcomments/controllers"
)

//SetLoginRouter Router para login
func SetLoginRouter(router *mux.Router)  {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}