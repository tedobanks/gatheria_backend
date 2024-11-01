package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"event_name" binding:"required"`
	Description string    `json:"event_description"`
	Location    string    `json:"event_location"`
	Date        time.Time `json:"event_date"`
	OrganizerId int64     `json:"event_organizer"`
	CreatedAt   time.Time `json:"created_at"`
}

func (event *Event) SaveEvent() error {
	query := `
    INSERT INTO event(name, description, location, date, organizer, created)
    VALUES (?, ?, ?, ?, ?, ?)
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("could not prepare query")

	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.OrganizerId, event.CreatedAt)
	if err != nil {
		panic("error countered executing")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Event created but unable to retrieve ID")
		return err
	}
	event.Id = lastInsertId
	fmt.Printf("Event created with ID:%v", lastInsertId)
	return nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM event"
	// Query is typically used to execute an SQL statement that returns Rows. So basically we can say Query is used for fetching data while Exec on the other hand is used to Execute a statement that does not return anything. For example INSERT, UPDATE OR DELETE. It is used to change things while Query is used to retrieve things
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// It is always essential to close rows as we don't want to leave them causing memory leaks. In this case i used the defer keyword as a safety net to close the row after the last row has been read
	defer rows.Close()

	// List to be populated and returned
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Date, &event.OrganizerId, &event.CreatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	var event Event
	query := "SELECT * FROM event WHERE id = ?"
	row := db.DB.QueryRow(query, eventId)

	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Date, &event.OrganizerId, &event.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
		UPDATE event
		SET name = ?, description = ?, location = ?, date = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.Id)
	return err

}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM event WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id)
	return err
}

func EmptyEventTable() error {
	query := "DELETE FROM event"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return err
}

func (event Event) RegisterForEvent(registrant_id int64) error {
	query := `
		INSERT INTO registrations(event_id, registrant_id)
		VALUES(?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id, registrant_id)
	return err
}
