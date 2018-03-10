package config



type StructParametros struct {
	Database map[string]string
	Modulo uint
}

	var  Parametros = StructParametros{
		Database:map[string]string{
			"host":"localhost",
			"port":"3306",
			"user":"root",
			"clave":"25448132",
			"database":"Proyecto2",

			},
			Modulo:1,
			}
