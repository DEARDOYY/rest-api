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
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	query := `	
	INSERT INTO events(name, description, localtion, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	e.ID = id

	// events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func main() {

}
