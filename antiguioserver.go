package main

import (
	"backend/conectarGeneral"
	"database/sql"
	"fmt"
	"net/http"
	"encoding/json"


)
type Mama struct{
	Camposnumericos map[string]string
	Melodia [12]int
}

type Mensajero struct{
	Mama
	Maneltol string
	OtraCosa string
}
 func (r Mama) MensajeEscrito (Json int ) (Resultado int){
 	fmt.Println("papafrita",Json)
 	Resultado = 50
 	return
 }
type Der interface {
	MensajeEscrito(int)int

}
func estoEs(dato Der){
	cd := dato.MensajeEscrito(1)
	fmt.Println("Esto es una Strut mensajero haciendo el metodo MensajeEscrito:",cd)
}
 func Hola(w http.ResponseWriter, r *http.Request) {
	var Menso = Mama{map[string]string{"manual1":"manual","manual2":"manualote",},[12]int{1,2,3,4,5,6,7,8,9,10,11,},}
	_=Menso.MensajeEscrito(12)
	var num  = [12]int{1,2,3,4,5,6,7,8,9,10,11,13}
	var structuraCompleja = Mensajero{
		Mama{
			map[string]string{
			"mamaes":"virginia",
			"papES":"nacho",
			},
			num,
		},
		"esto es un jeson de prueba",
		"y esto es otra cosa",
	}
	var dek Der = structuraCompleja
	estoEs(dek)
	json2, errw := json.Marshal(structuraCompleja)
	if errw != nil {
		fmt.Println("Algo salido mal y fue esto ", errw)
	}
	var j* []byte = &json2
	fmt.Println(*j)
	fmt.Println("esto es uno", r.Host)
	fmt.Println("esto es dos", r.UserAgent())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var ConexionDB *sql.DB
	ConexionDB = conectarGeneral.Crearpoll()
	fmt.Println("esto es tres", ConexionDB)
	var jesonmostrar *string = conectarGeneral.Traerusuario(ConexionDB)

	fmt.Fprint(w, string(*jesonmostrar))
	fmt.Fprintf(w,"<br>Esto es el valor dos:<br>")
	fmt.Fprint(w,string(*j))

}

func main() {
	mex := http.NewServeMux()



	mex.HandleFunc("/hola",Hola)
	public := http.FileServer(http.Dir("public"))
	mex.Handle("/", public)

	http.ListenAndServe(":90", mex)
}
