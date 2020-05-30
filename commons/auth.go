package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

var (
	privateKey *rsa.PrivateKey
	//PublicKey usada para validacion del token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("no lee el archivo privado->", err)
	}
	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("no lee el archivo publico->", err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("no se parseo llave privada->", err)
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("no se parseo llave publica", err)
	}
}

//GenetateJWT creacion de token
func GenetateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.hour * 2).Unix(),
			Issuer: "Escuela Digital",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	resul, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("no se firma token->", err)
	}
	return resul
}
