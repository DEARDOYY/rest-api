package models

import "example.com/rest-api/db"

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (? ,?)"

	// Prepare() + stmt.Exec() when we inserted data into the database
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Exec จะใช้กับ insert or update data ที่มีการ change data
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id

	return err
}
