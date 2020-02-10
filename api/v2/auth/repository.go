package auth

import (
	"database/sql"
)

type Repository interface {
	GetByEmail(email string) (User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) GetByEmail(email string) (User, error) {
	row := r.db.QueryRow("select id, role, email from users where email = ?", email)
	var u User
	err := row.Scan(&u.ID, &u.Role, &u.Email)
	if err != nil {
		return u, err
	}
	return u, nil
}

func NewRespository(db *sql.DB) Repository {
	return &repository{db}
}
