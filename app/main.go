package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/db"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/rest/routes"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	conn := db.NewDatabase()
	defer db.Close(conn)
	routes.UseRoutes(router, conn)

	log.Println("server started on port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("error starting server", err)
	}
}
