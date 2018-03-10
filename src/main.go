package src

import (
	"conectarGeneral"
	"github.com/gorilla/mux"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
)

type valores struct {
	Nombre string `json:"nombre"`
	Clave  string `json:"clave"`
}

func traerUsuario(w http.ResponseWriter, r *http.Request) {

	var valor valores
	fmt.Println("esto es uno", r.Host)
	fmt.Println("esto es dos", r.UserAgent())
	fmt.Println("esto es Authorization ", r.Header.Get("Authorization"))
	err := json.NewDecoder(r.Body).Decode(&valor)
	if err != nil {
		fmt.Println("algo salio mal")
	}
	fmt.Println("esto es body", valor)
	w.Header().Add("content-Type", "application/json; charset=utf-8; token=sjsycbsñeudbcebceubduebwñydgbscaudbwdusydywhqdwv")
	w.Header().Add("Etag", ` "499fd34e-29ec-42f695ca96761;48fe7523cfcc1"`)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	fmt.Println("esto es header", w.Header())
	var ConexionDB *sql.DB
	ConexionDB = conectarGeneral.Crearpoll()
	fmt.Println("esto es tres", ConexionDB)
	var jesonmostrar *[]byte = conectarGeneral.Traerusuario(ConexionDB)
	w.WriteHeader(http.StatusOK)
	w.Write(*jesonmostrar)

}
func verificarserver(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" || r.Method == "options" {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "authorization,content-type")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.WriteHeader(http.StatusOK)

	} else {
		traerUsuario(w, r)
	}
}

func main() {

	mex := mux.NewRouter().StrictSlash(true)
	go mex.HandleFunc("/hola", verificarserver)
	go mex.PathPrefix("/login").Handler(http.StripPrefix("/login", http.FileServer(http.Dir("public"))))
	go mex.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("public"))))

	//public := http.FileServer(http.Dir("public"))
	//mex.Handle("/", public)
	port := os.Getenv("PORT")
	if port == "" {
	port = "90"
	}

	myServer := &http.Server{
		Addr:           ":"+port,
		Handler:        mex,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	myServer.ListenAndServe()
}
