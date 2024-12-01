package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `	
	INSERT INTO events(name, description, localtion, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	// Prepare() + stmt.Exec() when we inserted data into the database
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	// Exec จะใช้กับ insert or update data ที่มีการ change data
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	// events = append(events, e)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	// query จะใช้กับการ get data ในรูปแบบ rows
	// DB.Query() when we fetched data from the database
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// ต้อง close เสมอเมื่อ functions ทำเสร็จ
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, err
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := "UPDATE events SET name = ?, description = ?, localtion = ?, date_time = ? WHERE id = ?"
	smst, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer smst.Close()
	_, err = smst.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	smst, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer smst.Close()
	_, err = smst.Exec(event.ID)

	return err
}
