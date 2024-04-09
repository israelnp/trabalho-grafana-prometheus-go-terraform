package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func UseRoutes(router *chi.Mux, dbConnection *sql.DB) {
	NewUserRoutes(router, dbConnection)
}
