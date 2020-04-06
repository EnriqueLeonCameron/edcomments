package routes

import (
	"github.com/gorilla/mux"
)

//es el que une todas las rutas es el que llama a cada una de las rutas configuradas
//inicia las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router) //metodos de archivos en este mismo paquete
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)

	return router
} 