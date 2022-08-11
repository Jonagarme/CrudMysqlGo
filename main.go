package main

import (
	"log"
	"net/http"

	"github.com/Jonagarme/CRUD-SQL/commons"
	"github.com/Jonagarme/CRUD-SQL/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()

	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose en el puerto 9000")

	log.Println(server.ListenAndServe())
}
