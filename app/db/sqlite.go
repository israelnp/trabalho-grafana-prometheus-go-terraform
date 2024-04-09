package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() *sql.DB {
	conn := connect()
	createTable(conn)
	return conn
}

func connect() *sql.DB {
	log.Println("connecting to the database...")
	conn, err := sql.Open("sqlite3", "./app.sqlite")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("connected to the database!")

	return conn
}

func Close(conn *sql.DB) {
	log.Println("closing the database connection...")
	err := conn.Close()
	if err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}
	log.Println("database connection closed!")
}

func createTable(conn *sql.DB) {
	log.Println("creating users table...")
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT,
		email TEXT NOT NULL UNIQUE,
		password TEXT
	)`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	log.Println("table created!")
}
