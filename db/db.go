package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Check if database connection is successful
	if err := DB.Ping(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Configure the database connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables() // Call your table creation function
}

// Call this function when the application is shutting down
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func createTables() {

	createUserTables := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NILL
	)`

	// DB.Exec()
	_, err := DB.Exec(createUserTables)
	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		localtion TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	// DB.Exec()
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
