package pet

import "database/sql"

type Pet struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Legs bool   `json:"legs"`
}

var db *sql.DB

func ConfigureDB(database *sql.DB) {
	db = database
}
