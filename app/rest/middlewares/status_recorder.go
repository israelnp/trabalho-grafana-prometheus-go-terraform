package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (statusRecorder *StatusRecorder) WriteHeader(statusCode int) {
	statusRecorder.statusCode = statusCode
	statusRecorder.ResponseWriter.WriteHeader(statusCode)
}

func GetRoutePattern(r *http.Request) string {
	reqContext := chi.RouteContext(r.Context())
	if pattern := reqContext.RoutePattern(); pattern != "" {
		return pattern
	}

	return "undefined"
}
