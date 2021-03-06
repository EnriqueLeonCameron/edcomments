package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/EnriqueLeonCameron/edcomments/controllers"
)

func SetCommentRouter(router *mux.Router)  {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")
	subRouter.HandleFunc("/", controllers.CommentGetAll ).Methods("GET")

	//para validar el token
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter), //este es el siguiente a ejecutarse que recibe ValidateToken
		),
	)
}