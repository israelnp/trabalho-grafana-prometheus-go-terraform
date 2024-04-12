package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

type PrometheusMiddleware struct {
	promService *services.PrometheusService
}

func NewPromService(promService *services.PrometheusService) *PrometheusMiddleware {
	return &PrometheusMiddleware{
		promService: promService,
	}
}

func (promMiddleware *PrometheusMiddleware) AddPrometheusMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		promMiddleware.promService.HttpRequestCounter.Inc()
		start := time.Now()

		statusRecorder := StatusRecorder{statusCode: 200, ResponseWriter: w}
		next.ServeHTTP(&statusRecorder, r)

		duration := time.Since(start)
		statusCode := strconv.Itoa(statusRecorder.statusCode)
		promMiddleware.promService.HttpStatusCounter.WithLabelValues(statusCode, r.Method).Inc()
		promMiddleware.promService.HttpRequestDuration.WithLabelValues(GetRoutePattern(r), r.Method, statusCode).Observe(duration.Seconds())
	})
}
