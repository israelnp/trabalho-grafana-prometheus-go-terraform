package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/models"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result, err := userHandler.UserService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error creating user: %v", err)
		return
	}
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userHandler.UserService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error listing users: %v", err)
		return
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
