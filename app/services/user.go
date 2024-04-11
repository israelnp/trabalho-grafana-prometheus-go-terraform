package services

import (
	"database/sql"
	"log"

	"github.com/israelnp/trabalho-grafana-prometheus-go-terraform/models"
)

type UserService struct {
	Conn *sql.DB
}

func NewUserService(dbConnection *sql.DB) *UserService {
	return &UserService{
		Conn: dbConnection,
	}
}

func (userService *UserService) CreateUser(user models.User) (*models.User, error) {
	log.Println("inserting user")

	query, err := userService.Conn.Prepare("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("error preparing query %v \n", err)
		return nil, err
	}

	result, err := query.Exec(nil, user.Name, user.Email, user.Password)
	if err != nil {
		log.Printf("error executing query %v \n", err)
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		log.Printf("error getting last insert id %v \n", err)
		return nil, err
	}

	defer query.Close()
	log.Println("successfully inserted user")
	return &models.User{
		ID:       userId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (userService *UserService) ListUsers() ([]models.User, error) {
	log.Println("listing users")

	rows, err := userService.Conn.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		log.Printf("error querying users %v \n", err)
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Printf("error scanning user %v \n", err)
			return nil, err
		}
		users = append(users, user)
	}

	defer rows.Close()
	log.Println("successfully listed users")
	return users, nil
}
