package data

import (
	"database/sql"
	"fmt"
)

type Seeder interface {
	Seed()
}

type seeder struct {
	db *sql.DB
}

const (
	usersTable = `
		drop table if exists users;
		create table users
		(
			id    INTEGER
				primary key autoincrement
				unique,
			role  integer default 1,
			email varchar(100),
			name  varchar(150)
		);
		INSERT INTO users (id, role, email, name) VALUES (1, 1, 'david@federnet.com', 'David (FederNet) Federspiel');
		INSERT INTO users (id, role, email, name) VALUES (2, 0, 'david.federspiel@ruralsourcing.com', 'David (RSI) Federspiel');`

	petsTable = `
		drop table if exists pets;
		create table pets
		(
			id    INTEGER
				primary key autoincrement
				unique,
			name  VARCHAR(50),
			age   INTEGER,
			legs  BOOLEAN,
			color VARCHAR(50)
		);
		INSERT INTO pets (id, name, age, legs, color) VALUES (1, 'Lonnie', 5, 1, 'Beige');
		INSERT INTO pets (id, name, age, legs, color) VALUES (2, 'Rick', 6, 1, 'Brown');
		INSERT INTO pets (id, name, age, legs, color) VALUES (3, 'Kibbles', 1, 1, 'Blue');
		INSERT INTO pets (id, name, age, legs, color) VALUES (4, 'Fred', 2, 1, 'Green');
		INSERT INTO pets (id, name, age, legs, color) VALUES (5, 'Slippers', 3, 0, 'Rainbow');
		INSERT INTO pets (id, name, age, legs, color) VALUES (6, 'Fuzzball', 12, 0, 'Spotted');
		INSERT INTO pets (id, name, age, legs, color) VALUES (7, 'Gary', 14, 1, 'Cream');`
)

func (s seeder) Seed() {
	_, err := s.db.Exec(petsTable)
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.db.Exec(usersTable)
	if err != nil {
		fmt.Println(err)
	}
}

func NewSeeder(db *sql.DB) Seeder {
	return &seeder{db}
}
