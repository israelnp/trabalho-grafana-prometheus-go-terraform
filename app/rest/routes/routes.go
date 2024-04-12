package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

func UseRoutes(router *chi.Mux, dbConnection *sql.DB, promService *services.PrometheusService) {
	NewUserRoutes(router, dbConnection)
}
