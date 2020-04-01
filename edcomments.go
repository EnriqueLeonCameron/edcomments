package main

import (
	"log"
	"flag"

	"github.com/EnriqueLeonCameron/edcomments/migration"
)

func main()  {
	//cuando se llame este archivo die me diante flag o "parametros", si quiero ejecutar la 
	//migracion(crear tablas en la base de datos) o no
	var migrate string

	//donde guardar,nombreFlag,valorDefecto,descripcion
	flag.StringVar(&migrate,"migrate","no","geenra la migracion a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migracion")
		migration.Migrate()
		log.Println("finalizó la migracion")
	}
}
//si coloco ./edcomments.exe --migration yes 
//hace la migracion
//si no le pongo el -migration lo toma como un no y no la hace