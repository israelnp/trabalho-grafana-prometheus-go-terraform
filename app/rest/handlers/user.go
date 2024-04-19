package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/models"
	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (userHandler *UserHandler) RandomStatusCodes(w http.ResponseWriter, r *http.Request) {
	statusCodes := []int{http.StatusOK, http.StatusInternalServerError, http.StatusBadRequest}
	randomStatusCode := rand.Intn(len(statusCodes))
	log.Printf("random status code index %d", randomStatusCode)
	time.Sleep(300 * time.Millisecond)

	log.Printf("returning status code %d", statusCodes[randomStatusCode])
	w.WriteHeader(statusCodes[randomStatusCode])
}

func (userHandler *UserHandler) SimulateDatabaseRead(w http.ResponseWriter, r *http.Request) {
	min := 1000
	timeOut := 3000
	randomMilliseconds := rand.Intn(5000-min+1) + min
	time.Sleep(time.Duration(randomMilliseconds) * time.Millisecond)

	if randomMilliseconds > timeOut {
		log.Printf("timeout reading database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("reading database in less than %d milliseconds", randomMilliseconds)
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) TakesRandomMiliseconds(w http.ResponseWriter, r *http.Request) {
	min := 100
	max := 1000
	randomMilliseconds := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(randomMilliseconds) * time.Millisecond)
	log.Printf("this request took only %d milliseconds", randomMilliseconds)
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	result, err := userHandler.userService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userHandler.userService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())
		log.Printf("Error listing users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
