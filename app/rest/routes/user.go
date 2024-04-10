package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/rest/handlers"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

func NewUserRoutes(router *chi.Mux, dbConnection *sql.DB, promService *services.PrometheusService) {
	userService := services.NewUserService(dbConnection)
	userHandler := handlers.NewUserHandler(userService, promService)

	router.Post("/users", userHandler.CreateUser)
	router.Get("/users", userHandler.ListUsers)
}
