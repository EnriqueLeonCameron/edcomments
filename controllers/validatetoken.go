package controllers

import (
	"net/http"
	"context"  //me permite enviar info del usuario entre los controladores
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"  //para ver lo que me viene en la peticion

	"github.com/EnriqueLeonCameron/edcomments/commons"
	"github.com/EnriqueLeonCameron/edcomments/models"
)

//ValidateToken validar eltoken del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc )  {
	//el tercer parametro es el siguinete controlador a ejecutar si el token es valido

	var m models.Message
	// este metodo request.ParseFromRequestWithClaims() me permite extraer info me mi reques "r"
	//request.OAuth2Extractor es el tipo de extraccion que utilizaré
	//el tercer parametro es un puntero a a estructura de mi token
	//el tercer parametro es una funcion anonima que devuelva la llave publica para validar si el token es valido o no 
	token,err := request.ParseFromRequestWithClaims(r, 
		request.OAuth2Extractor, 
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error){
			return commons.PublicKey, nil
		},
	)

	//err me puede decir si el token a expirado o si la fima no es valida, etc

	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) { //err.(type) me devuelve el tipo de error
			case *jwt.ValidationError:
				vError := err.(*jwt.ValidationError)
				switch vError.Errors{
					case jwt.ValidationErrorExpired:
						m.Message = "su token ha expirado"
						commons.DisplayMessage(w,m)
						return
					case jwt.ValidationErrorSignatureInvalid:
						m.Message = "La firma del token no coincide"
						commons.DisplayMessage(w,m)
						return
					default:
						m.Message = "Su token nno es valido"
						commons.DisplayMessage(w,m)
						return
				}
			
		}
	}
	if token.Valid{
		//para extraer los datos del usuario y no tener que volverlos a extraer en el sig controlador
		                            //A        //B    //C
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		//A:el contexto actual
		//b: como llamaremos a la llave para exttraer
		//C: De donde obtendremos la informacion
		next(w,r.WithContext(ctx))
	}else{
		m.Code = http.StatusUnauthorized
		m.Message = "Su token nno es valido"
		commons.DisplayMessage(w,m)
	}
}


