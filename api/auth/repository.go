package auth

import (
	"database/sql"
	"rsi.com/go-training/models"
)

type Repository interface {
	CurrentUser(email string) (models.User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) CurrentUser(email string) (models.User, error) {
	row := r.db.QueryRow("select id, role, email, name from users where email = ?", email)
	var u models.User
	err := row.Scan(&u.ID, &u.Role, &u.Email, &u.Name)
	if err != nil {
		return u, err
	}
	return u, nil
}

func NewRespository(db *sql.DB) Repository {
	return &repository{db}
}
