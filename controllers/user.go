package controllers

import(
	"log"
	"net/http"
	"fmt"
	"crypto/sha256"
	"crypto/md5"
	//"encoding/base64"
	"encoding/json"


	config "github.com/EnriqueLeonCameron/edcomments/configuration"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/EnriqueLeonCameron/edcomments/models"
	"github.com/EnriqueLeonCameron/edcomments/commons"
)


//login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)//el email y passwordse lo mapeo a user
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n",err)
		return
	}

	db :=config.GetConnection()
	defer db.Close()
			//user.password lo acceso así por que ya arriba le meti el email y password
	c := sha256.Sum256([]byte(user.Password))//para codificar el password que me llego en la peticion
							//en la variable c, lo que venga desde el inicio hasta la pos 32
	//pwd:= base64.URLEncoding.EncodeToString(c[:32]) //esto es lo mismo que  pwd := fmt.Sprintf("%x",c)
	//solo lo dejo por eje mplo de otra forma de hacerlo
	//mentira la camente por que no funciona, es mejor  pwd := fmt.Sprintf("%x",c)
	pwd := fmt.Sprintf("%x",c)

	//ahora tengo que comprara el password encriptado en mi bd y elq ue acaba de ingresar 
	db.Where("email = ? and password = ?", user.Email, pwd).First(&user) //El primer resultado que me venga se lo mape a user 
	if user.ID > 0{ //o sea si la consulta a la bd devolvio algo el id va ser mayor a cero
		user.Password = "" //para que el json lo ignore y no lo este enviando, por seguridad
		token := commons.GenerateJWT(user)
		
		j,err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}else{
		m := models.Message{
			Message: "Usuario o clave no validos",
			Code: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}

//permite registrar un usuario en la bd
func UserCreate(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s",err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = "Las Contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x",c)
	user.Password = pwd //la misma pero ahora codificada

	//para codificar la imagen, para las imagenes usaremos gravatar
	picmd5 := md5.Sum([]byte(user.Email)) //gravatar ocupa el correo codificado en md5
	picstr := fmt.Sprintf("%x",picmd5)

	//grabtar nos dará un enlace donde estara almacenado la picture 
														//?s=100 -> pixeles
	user.Picture = "https://gravatar.com/avatar/" + picstr +"?s=100"

	db := config.GetConnection()
	defer db.Close()

//.Error para que me capture un error en caso de que ocurra
	err = db.Create(&user).Error
	if err != nil {
		m.Message = "Error al guardar el usuario"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	m.Message = "Usuario creado con exito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w,m)
}
