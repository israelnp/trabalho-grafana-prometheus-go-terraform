package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/db"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/rest/middlewares"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/rest/routes"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	conn := db.NewDatabase()
	defer db.Close(conn)
	promService := services.NewPrometheusService()
	promMiddleware := middlewares.NewPromService(promService)
	router.Use(promMiddleware.AddPrometheusMiddleware)
	routes.UseRoutes(router, conn, promService)

	router.Handle("/metrics", promhttp.Handler())

	log.Println("server started on port 5000")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal("error starting server", err)
	}
}
