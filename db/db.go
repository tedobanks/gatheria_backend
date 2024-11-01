package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("could not open DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		firstname TEXT NOT NULL,
		lastname TEXT NOT NULL,
		isverified BOOLEAN NOT NULL,
		role TEXT NOT NULL,
		created DATETIME NOT NULL,
		updatedat DATETIME NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		organizer INTEGER NOT NULL,
		created DATETIME NOT NULL,
		FOREIGN KEY(organizer) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		registrant_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES event(id),
		FOREIGN KEY(registrant_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic(err)
	}

}
