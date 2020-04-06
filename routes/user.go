package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/EnriqueLeonCameron/edcomments/controllers"
)

/*
gorilla mux nos ayudara con las rutas
	"github.com/gorilla/mux"

negroni nos ayudará con los middlewares
	"github.com/urfave/negroni"
*/

//SetUSerRouter ruta para el registro de usuario
func SetUserRouter(router *mux.Router)  {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/",controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(negroni.Wrap(subRouter)),
	)

}