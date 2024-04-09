package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() *Database {
	conn := connect()
	createTable(conn)
	return &Database{
		Conn: conn,
	}
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
