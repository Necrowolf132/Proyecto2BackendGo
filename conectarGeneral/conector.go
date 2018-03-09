// Package conectarGeneral contiene funciones de utilidad para trabajar con cadenas html por http
package conectarGeneral

import (
	"backend/config"
	//"backend/nucleo/mysql-master"
	"backend/modelos"
	"database/sql"
	//"backend/nucleo"
	"encoding/json"
	"log"

)

var parametros *config.StructParametros = &config.Parametros
var parapoll string = parametros.Database["user"] + ":" + parametros.Database["clave"] + "@tcp(" + parametros.Database["host"] + ":" + parametros.Database["port"] + ")/" + parametros.Database["database"]

func Crearpoll() *sql.DB {

	ConexionBD, err := sql.Open("mysql", parapoll)
	if err != nil {
		panic(err)
	}
	return ConexionBD
}

func Traerusuario(ConexionDB *sql.DB) *[]byte {

	filas, err := ConexionDB.Query("SELECT * FROM `usuario`")

	if err != nil {
		log.Fatal(err)
	}
	defer filas.Close()
	var jsonEnvi []byte
	for filas.Next() {
		var Id_usuario int
		var Nombre string
		var Apellido string
		var Cedula int
		var Extra string
		var Id_perfil int
		var Id_seccion int
		if err := filas.Scan(&Id_usuario, &Nombre, &Apellido, &Cedula, &Extra, &Id_perfil, &Id_seccion); err != nil {
			log.Fatal(err)
		} else {


			var m = modelos.Usuario{Id_usuario, Nombre, Apellido, Cedula, Extra, Id_perfil, Id_seccion}
			jsonEnvi, err = json.Marshal(m)
			log.Println(string(jsonEnvi))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	if err := filas.Err(); err != nil {
		log.Fatal(err)
	}

	return &jsonEnvi
}
