package main

import (
	"log"
	"flag"
	"net/http"
	
	"github.com/urfave/negroni"
	
	"github.com/EnriqueLeonCameron/edcomments/migration"
	"github.com/EnriqueLeonCameron/edcomments/routes"
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

	//inicia las rutas
	router := routes.InitRoutes()

	//inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//inicia el servidor
	server := &http.Server{
		Addr: ":8080",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecucion del programa")
}
//si coloco ./edcomments.exe --migration yes 
//hace la migracion
//si no le pongo el -migration lo toma como un no y no la hace

