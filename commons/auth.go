package commons

//aqui es donde firmaré los tokens para ser validados

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/EnriqueLeonCameron/edcomments/models"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
)

func init()  {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("no se pudo hacer parse a privateKey")
	}
	
}


//Genera EL TOKEN PARA EL CLIENTE
func GenerateJWT(user models.User) string  {
	claims := models.Claim{
		 User: user,
		 StandardClaims: jwt.StandardClaims{
			//Para que la sesion expire en 2 horas
			//ExpirateAt: time.now().Add(time.Hour * 2).Unix(),
			Issuer: "Escuela Digital",
		 },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result,err := token.SignedString(privateKey) //aqui firmo el token
	if err != nil {
		log.Fatal("no se pudo firmar el token")
	}

	return result
}