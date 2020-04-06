package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EnriqueLeonCameron/edcomments/models"
	"github.com/EnriqueLeonCameron/edcomments/configuration"
	"github.com/EnriqueLeonCameron/edcomments/commons"
)

//controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request){
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el voto a registrar: %s", err) 
		commons.DisplayMessage(w,m)
		return
	}
	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	//para ver si el usuario ya voto por el comentario que está tratando de votar, se va guardar en currentVote
	db.Where("comment_id = ? and user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)

	//si no existe es por que no ha votado todavia
	if currentVote.ID == 0 {
		db.Create(&vote) //estonces creo el voto nuevo
		err := updateCommentVotes(vote.CommentID, vote.Value, false) // y actualizo el total de votos
												//pasandole el id del comentario y el valor del voto
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w,m)  
			return
		}
		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w,m)  
		return
	}else if currentVote.Value != vote.Value{ //si esta votando direfente al que tenia, o se ael contrario
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value, true)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w,m) 
			return
		}
		m.Message = "Voto actualizado" 
		m.Code = http.StatusOK
		commons.DisplayMessage(w,m) 
		return
	}
	m.Message = "Este voto ya esta registrado" 
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w,m) 
}

//Actualiza la cantidad de votos en la tabla comentarios, iupdate indica si es un voto para actualziar
func updateCommentVotes(commentID uint, vote bool, isUpdate bool) (err error){
	comment := models.Comment{}

	db := configuration.GetConnection()
	defer db.Close()

	//obtengo el comentario
	rows := db.First(&comment, commentID).RowsAffected //el id por el que filtro es commentID y se guardará en comment
	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate{
				comment.Votes++
			}
		}else{
			comment.Votes--
			if isUpdate{
				comment.Votes--
			}
		}
		db.Save(&comment)//Se actualiza en la bd
	}else{
		err = errors.New("no se encontro un registro de comentario para signar el voto")
	}
	return
	
}