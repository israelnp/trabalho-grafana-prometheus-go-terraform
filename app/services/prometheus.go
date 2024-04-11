package services

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PrometheusService struct {
	HttpRequestCounter  prometheus.Counter
	HttpStatusCounter   prometheus.CounterVec
	HttpRequestDuration prometheus.HistogramVec
}

func NewPrometheusService() *PrometheusService {
	return &PrometheusService{
		HttpRequestCounter:  createHttpRequestCounter(),
		HttpStatusCounter:   createHttpStatusCounter(),
		HttpRequestDuration: createHttpRequestDuration(),
	}
}

func createHttpRequestCounter() prometheus.Counter {
	httpRequestCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_request_counter_total",
		Help: "The total number of HTTP requests",
	})

	return httpRequestCounter
}

func createHttpStatusCounter() prometheus.CounterVec {
	httpStatusCounter := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_status_counter_total",
		Help: "The total number of HTTP requests by status",
	}, []string{"code", "method"})

	httpStatusCounter.WithLabelValues("200", "GET").Add(0)
	httpStatusCounter.WithLabelValues("200", "POST").Add(0)
	httpStatusCounter.WithLabelValues("500", "GET").Add(0)
	httpStatusCounter.WithLabelValues("500", "POST").Add(0)

	return *httpStatusCounter
}

func createHttpRequestDuration() prometheus.HistogramVec {
	httpRequestDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "The duration of HTTP requests",
	}, []string{"route", "method", "status"})

	return *httpRequestDuration
}
