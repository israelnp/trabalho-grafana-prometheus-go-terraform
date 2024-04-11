package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/models"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

type UserHandler struct {
	UserService *services.UserService
	PromService *services.PrometheusService
}

func NewUserHandler(userService *services.UserService, promService *services.PrometheusService) *UserHandler {
	return &UserHandler{
		UserService: userService,
		PromService: promService,
	}
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	userHandler.PromService.HttpRequestCounter.Inc()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result, err := userHandler.UserService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		durantion := time.Since(start)
		userHandler.PromService.HttpRequestDuration.WithLabelValues("/users", "POST", "500").Observe(durantion.Seconds())
		userHandler.PromService.HttpStatusCounter.WithLabelValues("500", "POST").Inc()
		return
	}

	durantion := time.Since(start)
	userHandler.PromService.HttpRequestDuration.WithLabelValues("/users", "POST", "200").Observe(durantion.Seconds())
	userHandler.PromService.HttpStatusCounter.WithLabelValues("200", "POST").Inc()
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	userHandler.PromService.HttpRequestCounter.Inc()
	users, err := userHandler.UserService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error listing users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		durantion := time.Since(start)
		userHandler.PromService.HttpRequestDuration.WithLabelValues("/users", "GET", "500").Observe(durantion.Seconds())
		userHandler.PromService.HttpStatusCounter.WithLabelValues("500", "GET").Inc()
		return
	}

	durantion := time.Since(start)
	userHandler.PromService.HttpRequestDuration.WithLabelValues("/users", "GET", "200").Observe(durantion.Seconds())
	userHandler.PromService.HttpStatusCounter.WithLabelValues("200", "GET").Inc()
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
