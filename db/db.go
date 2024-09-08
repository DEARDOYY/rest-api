package db

import "database/sql"

var DB *sql.DB

func initDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database. ")
	}

	// pool connection
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
