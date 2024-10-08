package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	err = DB.Ping()
	if err != nil {
		panic("Error connecting to the database")
	}

	if DB == nil {
		panic("Database is not initialized")
	}

	// pool connection
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

	defer DB.Close()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		localtion TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
