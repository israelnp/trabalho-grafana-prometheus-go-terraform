package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/rest/handlers"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

func NewUserRoutes(router *chi.Mux, dbConnection *sql.DB) {
	userService := services.NewUserService(dbConnection)
	userHandler := handlers.NewUserHandler(userService)

	router.Post("/users", userHandler.CreateUser)
	router.Get("/users", userHandler.ListUsers)
	router.Get("/bad_data_base_read", userHandler.SimulateDatabaseRead)
	router.Get("/random_miliseconds", userHandler.TakesRandomMiliseconds)
	router.Get("/random_status_codes", userHandler.RandomStatusCodes)
}
