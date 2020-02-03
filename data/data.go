package data

import (
	"database/sql"
)

var db *sql.DB

func ConfigureDB(database *sql.DB) {
	db = database
}
