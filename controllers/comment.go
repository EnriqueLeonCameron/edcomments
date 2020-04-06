package controllers

import (
	"strconv"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	//jwt "github.com/dgrijalva/jwt-go"

	"github.com/EnriqueLeonCameron/edcomments/models"
	"github.com/EnriqueLeonCameron/edcomments/configuration"
	"github.com/EnriqueLeonCameron/edcomments/commons"
)

//CommentCreate permite registrar un comentario
func CommentCreate(w http.ResponseWriter, r *http.Request)  {
	comment := models.Comment{}
	user := models.User{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(&comment)  //json.NewDecoder() me pasa algo a json, en la estructura que yo le diga
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w,m)
		return
	}

	comment.UserID = user.ID 

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error //aqui creo el comentario
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario: %s", err)
		commons.DisplayMessage(w,m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con exito"
	commons.DisplayMessage(w,m)
	
}


//CommentGetAll obtiene todos los comentarios
func CommentGetAll(w http.ResponseWriter, r *http.Request){
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}
							//esto .(models.User) e un castin
	user, _ = r.Context().Value("user").(models.User) //Recerde que ya el usuario viene en el contexto del request
	vars := r.URL.Query()
	//la ruta puede ser:  /api/comments/?order=votes&idlimit=10
	//con r.URL.Query() obtengo los parametros, o sea todo despues del ?

	db := configuration.GetConnection()
	defer db.Close()

	cComment := db.Where("parent_id = 0")
	//la llave order puede estar varias veces, por lo tanto se guarda como un slice, pero yo tomo la primera
	if order,ok := vars["order"]; ok {//si existe la llave order, ok va ser verdadero y el valor de la 
		if order[0]	== "votes"{		//llave va quedar guradado en la variable order
			cComment = cComment.Order("votes desc, created_at desc")
		}
	}else{
		if idlimit, ok := vars["idlimit"]; ok{
			registerByPage := 30
			//offset desde donde quede la ultima vez para buscar mas
			//strconv.Atoi() para voncertir un string a int
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("Error:", err)
			}
			cComment = cComment.Where("id BETWEEN ? AND ?", offset-registerByPage, offset )
		}
		cComment = cComment.Order("id desc")
	}

	cComment.Find(&comments)//lo que haya en cComments busquelo y pongalo en el slice comments

	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User) //es como una relacion
		comments[i].User[0].Password = ""// lo dejo vacio para que el password no viaje por el json 
		comments[i].Children = commentGetChildren(comments[i].ID)
		
		//se busca el voto de usuario en sesion
		vote.CommentID = comments[i].ID
		vote.UserID = user.ID
		count := db.Where(&vote).Find(&vote).RowsAffected
		if count > 0 {//si es mayor que cero es por que el usuario voto en el comentario
			if vote.Value {
				comments[i].HasVote = 1
			}else{
				comments[i].HasVote = -1
			}
		}
	}

	j, err := json.Marshal(comments)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message= "Error al convertir los comentariosen json"
		commons.DisplayMessage(w,m)
		return
	}

	if len(comments) > 0{ //si hay contenido
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}else{ //si no hay contenido
		m.Code = http.StatusNoContent
		m.Message = "No se encontraron comentarios"
		commons.DisplayMessage(w,m)
	}
	//hasta aqui muestra todos los comentarios
	//falta saber si el usuario que hace la consilta de los comentarios ha votado en ellos
}

func commentGetChildren(id uint) (children []models.Comment) {
	db := configuration.GetConnection()
	defer db.Close()

	db.Where("parent_id = ?", id).Find(&children)
	for i := range children {
		db.Model(&children[i]).Related(&children[i].User) //es como una relacion
		children[i].User[0].Password = ""// lo dejo vacio para que el password no viaje por el json 
		
	}
	return  //ya dije arriba que retornaba por eso solo pongo return
}