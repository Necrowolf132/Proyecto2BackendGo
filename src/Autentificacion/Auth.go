package Autentificacion

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"encoding/json"
	"crypto/rsa"
	"modelos"
	"net/http"
	"io/ioutil"
	"log"
	"time"
)


var PrivadaKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func init(){
	 BytesPrivat ,err := ioutil.ReadFile("../keys/private.rsa")
	 if err != nil {
	 	log.Println("Ocurrio un error con la lectura de la carga Keys privada")

	 }
	 BytesPublic, err := ioutil.ReadFile("../keys/public.rsa.pub")
	if err != nil {
		log.Println("Ocurrio un error con la lectura de la carga Keys publica")

	}
	PrivadaKey, err = jwt.ParseRSAPrivateKeyFromPEM(BytesPrivat)
	if err != nil {
		log.Println("Ocurrio un error con el parse de la lectura Keys Private keys")

	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(BytesPublic)
	if err != nil {
		log.Println("Ocurrio un error con el parse de la lectura Keys Publica keys")

	}
}

func crearTOken (user modelos.Usuario)(string){
		var EstructuraToken modelos.Token = modelos.Token{
			user.Id_usuario,
			user.Id_perfil,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute *15).Unix(),
				Issuer: "confirmacion de usuario",
			},
		}
		Token := jwt.NewWithClaims(jwt.SigningMethodRS256,EstructuraToken)
		TokenFinal, err := Token.SignedString(PrivadaKey)
		if(err != nil ){
			log.Println("Ocurrio un error al firmar el token")
		}
		return TokenFinal
}
func VerificarToken (r *http.Request)  *[]byte  {
	var tokenVerificado *jwt.Token
	var RespondeValido modelos.ModeloRespuestaToken
	var err error
	tokenVerificado, err = request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &modelos.Token{},
		func(token *jwt.Token) (interface{},error){
			return  PublicKey,nil
		})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			verError := err.(*jwt.ValidationError)
			switch verError.Errors {
			case jwt.ValidationErrorExpired:
				log.Println("El tiempo del token expiro")
				RespondeValido = modelos.ModeloRespuestaToken{
					"El tiempo del token expiro",
					false,
				}
			case jwt.ValidationErrorSignatureInvalid:
				log.Println("La firma del toquen no coincide")
				RespondeValido = modelos.ModeloRespuestaToken{
					"La firma del toquen no coincide",
					false,
				}
			default:
				log.Println("El token no es valido")
				RespondeValido = modelos.ModeloRespuestaToken{
					"El token no es valido",
					false,
				}
			}
		default:
			log.Println("El token no es valido")
			RespondeValido = modelos.ModeloRespuestaToken{
				"El token no es valido",
				false,
			}
		}
	}
	if tokenVerificado.Valid {
		log.Println("El token a sido autenticado")
		RespondeValido = modelos.ModeloRespuestaToken{
			"El token a sido autenticado",
			true,
		}
	}
	var jsonRespondeValido []byte
	jsonRespondeValido, err = json.Marshal(RespondeValido)
	return &jsonRespondeValido

}