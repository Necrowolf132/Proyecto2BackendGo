package importjson
import(
	"os"
	"encoding/json"
	"fmt"
)



func LeerJson(ruta string) *interface{}{
	Abierto,_ := os.Open(ruta)
	var contenido = make([]byte,120)
	_, er := Abierto.Read(contenido)
	if(er!=nil){
		fmt.Println(er)
		}
	defer Abierto.Close()
	var jsontal interface{}
	er = json.Unmarshal(contenido,&jsontal)
	if(er!=nil){
		fmt.Println(er)
	}
	return &jsontal

}
