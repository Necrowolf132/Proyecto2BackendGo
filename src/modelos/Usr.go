package modelos
import jwt "github.com/dgrijalva/jwt-go"
type Usuario struct {
	Id_usuario int    `json:"id_usuario"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Cedula     int    `json:"cedula"`
	Extra      string `json:"extra"`
	Id_perfil  int    `json:"id_perfil"`
	Id_seccion int    `json:"id_seccion"`

}

type Token struct {
	Id_user int `json:"id_usuario"`
	Id_perfil int  `json:"id_perfil"`
	jwt.StandardClaims
}